package template

import (
    "errors"
    "fmt"
    "io"
    "io/fs"
    "log"
    "path/filepath"
    "strings"

    "github.com/nikolalohinski/gonja"
    "github.com/nikolalohinski/gonja/exec"
    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/asset"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
)

type TemplateRenderer struct {
    applicationConfig config.ApplicationConfig
    assetRenderer *asset.AssetRenderer
    gonjaEnvironment *gonja.Environment
    rootDir config.TemplatesDirectory
}

func NewTemplateRenderer(
    applicationConfig config.ApplicationConfig,
    assetRenderer *asset.AssetRenderer,
    gonjaEnvironment *gonja.Environment,
    rootDir config.TemplatesDirectory,
) *TemplateRenderer {
    return &TemplateRenderer{applicationConfig, assetRenderer, gonjaEnvironment, rootDir}
}

func (templateRenderer *TemplateRenderer) Load() error {
    rootDir := string(templateRenderer.rootDir)
    walker := func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if d.IsDir() {
            return nil
        }

        relPath, err := filepath.Rel(rootDir, path)
        if err != nil {
            return err
        }

        if strings.HasSuffix(relPath, ".swp") {
            return nil
        }

        _, cErr := templateRenderer.gonjaEnvironment.FromCache(relPath)
        if cErr != nil {
            log.Printf("Failed to cache template %s", relPath)
            return cErr
        }

        return err
    }
    return filepath.WalkDir(rootDir, walker)
}

func (templateRenderer *TemplateRenderer) Render(writer io.Writer, name string, data interface{}, _ ...string) error {
    template, err := templateRenderer.gonjaEnvironment.FromCache(name)
    if err != nil {
        return err
    }

    viewContext, err := templateRenderer.populateViewContext(data)
    if err != nil {
        return err
    }

    result, err := template.ExecuteBytes(viewContext)
    if err != nil {
        log.Fatal(err)
        return err
    }

    writer.Write(result)

    return nil
}

func (templateRenderer *TemplateRenderer) RenderToString(name string, data interface{}) (string, error) {
    template, err := templateRenderer.gonjaEnvironment.FromCache(name)
    if err != nil {
        return "", err
    }

    viewContext, err := templateRenderer.populateViewContext(data)
    if err != nil {
        return "", err
    }

    result, err := template.Execute(viewContext)
    if err != nil {
        log.Fatal(err)
        return result, err
    }

    return result, nil
}

func (templateRenderer *TemplateRenderer) populateViewContext(data interface{}) (map[string]interface{}, error) {
    viewContext, isMap := data.(map[string]interface{})
    if !isMap {
        return viewContext, errors.New("jinja view data must be a map")
    }

    message := "fiber.Ctx not found in template context"
    _, exists := viewContext["ctx"]
    if !exists {
        return viewContext, errors.New(message)
    }

    ctx, ok := viewContext["ctx"].(*fiber.Ctx)
    if !ok {
        return viewContext, errors.New(message)
    }

    viewContext["request_context"] = ctx
    viewContext["get_asset_html"] = func(va *exec.VarArgs) *exec.Value {
        filename := va.Args[0].String()
        assetType := va.Args[1].String()

        result := templateRenderer.assetRenderer.GetAssetHtml(ctx, filename, assetType)
        return exec.AsSafeValue(result)
    }
    viewContext["route"] = func(va *exec.VarArgs) *exec.Value {
        routeName := va.Args[0].String()
        routeParameters := make(map[string]interface{})
        if len(va.Args) == 2 {
            if !va.Args[1].IsDict() {
                return exec.AsValue(errors.New("route parameters must be a map"))
            }

            dict, ok := va.Args[1].Interface().(*exec.Dict)
            if ok {
                for _, key := range dict.Keys() {
                    strKey, strOk := key.Interface().(string)
                    if strOk {
                        routeParameters[strKey] = dict.Get(key).Interface()
                    }
                }
            }
        }
        result, err := ctx.GetRouteURL(routeName, fiber.Map(routeParameters))
        if err != nil {
            return exec.AsValue(err)
        }
        return exec.AsSafeValue(result)
    }
    viewContext["icon"] = func(va *exec.VarArgs) *exec.Value {
        name := va.Args[0].String()
        className := va.Args[1].String()

        var label string
        if len(va.Args) == 3 {
            label = va.Args[2].String()
        }

        var ariaAttribute string
        if label == "" {
            ariaAttribute = "aria-hidden=\"true\""
        } else {
            ariaAttribute = fmt.Sprintf("aria-label=\"%s\"", label)
        }

        result := fmt.Sprintf(
            "<svg class=\"%s\" %s><use xlink:href=\"#icon-sprite__%s\"/></svg>",
            className,
            ariaAttribute,
            name,
        )
        return exec.AsSafeValue(result)
    }
    viewContext["config"] = templateRenderer.applicationConfig

    return viewContext, nil
}
