package template

import (
    "errors"
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

    viewContext["get_asset_html"] = BuildAssetFunction(templateRenderer.assetRenderer, ctx)
    viewContext["route"] = BuildRouteFunc(ctx)

    return viewContext, nil
}

type GonjaFunc func(*exec.VarArgs) *exec.Value

func BuildAssetFunction(assetRenderer *asset.AssetRenderer, ctx *fiber.Ctx) GonjaFunc {
    return func(va *exec.VarArgs) *exec.Value {
        args := va.ExpectArgs(2)

        if !args.Args[0].IsString() {
            return exec.AsValue(errors.New("first argument to “get_asset_html” function should be a string"))
        }

        if !args.Args[1].IsString() {
            return exec.AsValue(errors.New("second argument to “get_asset_html” function should be a string"))
        }

        filename := args.Args[0].String()
        assetType := args.Args[1].String()

        result := assetRenderer.GetAssetHtml(ctx, filename, assetType)
        return exec.AsSafeValue(result)
    }
}

func BuildRouteFunc(ctx *fiber.Ctx) GonjaFunc {
    return func(va *exec.VarArgs) *exec.Value {
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
        return exec.AsValue(result)
    }
}
