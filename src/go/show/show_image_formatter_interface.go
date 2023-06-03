package show

import (
    "context"
)

type ShowImageResult struct {
    Default string
    Srcset string
}

type ShowImageFormatterInterface interface {
    Format(ctx context.Context, imageType string, imagePath string, defaultSize string) ShowImageResult
    FormatUrl(ctx context.Context, imagePath string, size string) string
}
