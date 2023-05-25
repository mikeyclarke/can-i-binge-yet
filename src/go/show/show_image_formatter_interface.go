package show

import (
    "context"
)

type ShowImageFormatterInterface interface {
    Format(ctx context.Context, imageType string, imagePath string, defaultSize string) ShowImageResult
}
