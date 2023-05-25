package asset

import (
    "encoding/json"
    "errors"
    "os"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
)

type AssetManifest struct {
    manifestPath config.AssetManifestPath
}

func NewAssetManifest(manifestPath config.AssetManifestPath) *AssetManifest {
    return &AssetManifest{manifestPath}
}

func (assetManifest *AssetManifest) GetAssetUrl(filename string) string {
    contents, err := os.ReadFile(string(assetManifest.manifestPath))
    if err != nil {
        panic(err)
    }

    data := map[string]string{}
    json.Unmarshal(contents, &data)

    url, exists := data[filename]
    if !exists {
        panic(errors.New("asset not found"))
    }

    return url
}
