package asset

import (
    "errors"
    "fmt"

    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/file"
    idGenerator "github.com/mikeyclarke/can-i-binge-yet/src/go/id_generator"
)

type AssetRenderer struct {
    assetDirectory config.AssetDirectory
    assetManifest AssetManifestInterface
    fileReader file.FileReaderInterface
    idGenerator idGenerator.AlphanumericIdGeneratorInterface
}

func NewAssetRenderer(
    assetDirectory config.AssetDirectory,
    assetManifest AssetManifestInterface,
    fileReader file.FileReaderInterface,
    idGenerator idGenerator.AlphanumericIdGeneratorInterface,
) *AssetRenderer {
    return &AssetRenderer{assetDirectory, assetManifest, fileReader, idGenerator}
}

func (assetRenderer *AssetRenderer) GetAssetHtml(ctx *fiber.Ctx, filename string, assetType string) string {
    assetDigest := NewAssetDigestFromRequestHeaders(ctx.GetReqHeaders())
    assetUrl := assetRenderer.assetManifest.GetAssetUrl(filename)

    if assetDigest.Has(filename) && assetDigest.Get(filename) == assetUrl {
        result, err := formatExternalResourceHtml(assetUrl, assetType)
        if err != nil {
            panic(err)
        }
        return result
    }

    assetPath := fmt.Sprintf("%s%s", assetRenderer.assetDirectory, assetUrl)

    content, err := assetRenderer.fileReader.Read(assetPath)
    if err != nil {
        panic(err)
    }

    result, err := formatInlineHtml(ctx, assetRenderer.idGenerator, filename, assetUrl, content, assetType)
    if err != nil {
        panic(err)
    }
    return result
}

func formatInlineHtml(
    ctx *fiber.Ctx,
    idGenerator idGenerator.AlphanumericIdGeneratorInterface,
    filename string,
    url string,
    content string,
    assetType string,
) (string, error) {
    result := ""
    err := error(nil)

    switch assetType {
        case "css":
            stylesheetId := idGenerator.Generate(8)

            cacheableStylesheets := ctx.Locals("cacheable_stylesheets")
            cacheableStylesheetsMap, ok := cacheableStylesheets.(map[string]map[string]string)
            if !ok || nil == cacheableStylesheetsMap {
                cacheableStylesheetsMap = make(map[string]map[string]string)
            }

            cacheableStylesheetsMap[stylesheetId] = map[string]string {
                "bundle": filename,
                "src": url,
            }

            ctx.Locals("cacheable_stylesheets", cacheableStylesheetsMap)

            result = fmt.Sprintf("<style id=\"%s\">%s</style>", stylesheetId, content)
        case "js":
            result = fmt.Sprintf(
                "<cacheable-asset src=\"%s\" bundle=\"%s\"><script>%s</script></cacheable-asset>",
                url,
                filename,
                content,
            )
        default:
            err = errors.New("unsupported asset type")
    }

    return result, err
}

func formatExternalResourceHtml(url string, assetType string) (string, error) {
    result := ""
    err := error(nil)

    switch assetType {
        case "css":
            result = fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\">", url)
        case "js":
            result = fmt.Sprintf("<script src=\"%s\"></script>", url)
        default:
            err = errors.New("unsupported asset type")
    }

    return result, err
}
