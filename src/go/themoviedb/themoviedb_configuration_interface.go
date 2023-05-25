package themoviedb

import (
    "context"
)

type TheMovieDbConfigurationInterface interface {
    GetImageBaseUrl(ctx context.Context) string
    GetImageSizes(ctx context.Context, imageType string) []string
}
