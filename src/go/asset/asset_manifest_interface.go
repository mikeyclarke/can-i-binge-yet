package asset

type AssetManifestInterface interface {
    GetAssetUrl(filename string) string
}
