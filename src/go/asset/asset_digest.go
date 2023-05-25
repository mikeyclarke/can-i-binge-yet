package asset

import (
    "encoding/json"
)

type AssetDigest struct {
    digest map[string]string
}

func NewAssetDigestFromRequestHeaders(requestHeaders map[string]string) *AssetDigest {
    var digest map[string]string
    value, exists := requestHeaders["X-Asset-Digest"]
    if !exists || value == "" {
        digest = make(map[string]string)
    } else {
        digest = map[string]string{}
        json.Unmarshal([]byte(value), &digest)
    }

    return &AssetDigest{digest}
}

func (assetDigest *AssetDigest) Has(filename string) bool {
    _, exists := assetDigest.digest[filename]
    return exists
}

func (assetDigest *AssetDigest) Get(filename string) string {
    return assetDigest.digest[filename]
}
