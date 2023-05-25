package asset

import (
    "fmt"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/valyala/fasthttp"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/file"
    idGenerator "github.com/mikeyclarke/can-i-binge-yet/src/go/id_generator"
)

type AssetRendererTestSuite struct {
    suite.Suite
    assetDirectory config.AssetDirectory
    assetManifest *MockAssetManifestInterface
    fileReader *file.MockFileReaderInterface
    idGenerator *idGenerator.MockAlphanumericIdGeneratorInterface
    assetRenderer *AssetRenderer
    ctx *fiber.Ctx
}

func (suite *AssetRendererTestSuite) SetupTest() {
    suite.assetDirectory = config.AssetDirectory("public")
    suite.assetManifest = NewMockAssetManifestInterface(suite.T())
    suite.fileReader = file.NewMockFileReaderInterface(suite.T())
    suite.idGenerator = idGenerator.NewMockAlphanumericIdGeneratorInterface(suite.T())
    suite.assetRenderer = NewAssetRenderer(
        suite.assetDirectory,
        suite.assetManifest,
        suite.fileReader,
        suite.idGenerator,
    )

    requestCtx := fasthttp.RequestCtx{}
    app := fiber.New()
    suite.ctx = app.AcquireCtx(&requestCtx)
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithCssAssetInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"app.css\": \"/compiled/app.abcd1234.css\"}")
    filename := "app.css"
    assetType := "css"

    url := "/compiled/app.abcd1234.css"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)

    expected := fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\">", url)
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithCssAssetNotInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"foo.css\": \"/compiled/foo.abcd1234.css\"}")
    filename := "app.css"
    assetType := "css"

    url := "/compiled/app.abcd1234.css"
    assetPath := fmt.Sprintf("%s%s", suite.assetDirectory, url)
    content := ".foo { color: red; }"
    stylesheetId := "h4mEg08q"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)
    suite.fileReader.On("Read", assetPath).Return(content, nil)
    suite.idGenerator.On("Generate", 8).Return(stylesheetId)

    expectedLocal := map[string]map[string]string{
        stylesheetId: map[string]string{
            "bundle": filename,
            "src": url,
        },
    }

    expected := fmt.Sprintf("<style id=\"%s\">%s</style>", stylesheetId, content)
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
    assert.Equal(suite.T(), expectedLocal, suite.ctx.Locals("cacheable_stylesheets"))
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithOutdatedCssAssetInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"app.css\": \"/compiled/app.abcd1234.css\"}")
    filename := "app.css"
    assetType := "css"

    url := "/compiled/app.efgh5678.css"
    assetPath := fmt.Sprintf("%s%s", suite.assetDirectory, url)
    content := ".foo { color: red; }"
    stylesheetId := "h4mEg08q"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)
    suite.fileReader.On("Read", assetPath).Return(content, nil)
    suite.idGenerator.On("Generate", 8).Return(stylesheetId)

    expectedLocal := map[string]map[string]string{
        stylesheetId: map[string]string{
            "bundle": filename,
            "src": url,
        },
    }

    expected := fmt.Sprintf("<style id=\"%s\">%s</style>", stylesheetId, content)
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
    assert.Equal(suite.T(), expectedLocal, suite.ctx.Locals("cacheable_stylesheets"))
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithJsAssetInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"app.js\": \"/compiled/app.abcd1234.js\"}")
    filename := "app.js"
    assetType := "js"

    url := "/compiled/app.abcd1234.js"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)

    expected := fmt.Sprintf("<script src=\"%s\"></script>", url)
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithJsAssetNotInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"foo.js\": \"/compiled/foo.abcd1234.js\"}")
    filename := "app.js"
    assetType := "js"

    url := "/compiled/app.abcd1234.js"
    assetPath := fmt.Sprintf("%s%s", suite.assetDirectory, url)
    content := "alert('Foo');"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)
    suite.fileReader.On("Read", assetPath).Return(content, nil)

    expected := fmt.Sprintf(
        "<cacheable-asset src=\"%s\" bundle=\"%s\"><script>%s</script></cacheable-asset>",
        url,
        filename,
        content,
    )
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
}

func (suite *AssetRendererTestSuite) TestGetAssetHtmlWithOutdatedJsAssetInRequestHeader() {
    suite.ctx.Context().Request.Header.Add("X-Asset-Digest", "{\"app.js\": \"/compiled/app.abcd1234.js\"}")
    filename := "app.js"
    assetType := "js"

    url := "/compiled/app.efgh5678.js"
    assetPath := fmt.Sprintf("%s%s", suite.assetDirectory, url)
    content := "alert('Foo');"

    suite.assetManifest.On("GetAssetUrl", filename).Return(url)
    suite.fileReader.On("Read", assetPath).Return(content, nil)

    expected := fmt.Sprintf(
        "<cacheable-asset src=\"%s\" bundle=\"%s\"><script>%s</script></cacheable-asset>",
        url,
        filename,
        content,
    )
    result := suite.assetRenderer.GetAssetHtml(suite.ctx, filename, assetType)

    assert.Equal(suite.T(), expected, result)
}

func TestAssetRendererTestSuite(t *testing.T) {
    suite.Run(t, new(AssetRendererTestSuite))
}
