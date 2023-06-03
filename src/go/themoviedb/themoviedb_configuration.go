package themoviedb

import (
    "context"
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/cache"
)

const (
    CacheKey = "tmdb_configuration"
    CacheLifetime = time.Hour * 120 // 5 days
)

type TheMovieDbConfiguration struct {
    cache cache.CacheInterface
    client *TheMovieDbClient
}

func NewTheMovieDbConfiguration(cache cache.CacheInterface, tmdbClient *TheMovieDbClient) *TheMovieDbConfiguration {
    return &TheMovieDbConfiguration{cache, tmdbClient}
}

func (tmdbConfiguration *TheMovieDbConfiguration) GetImageBaseUrl(ctx context.Context) string {
    configuration := tmdbConfiguration.getConfiguration(ctx)
    return configuration.Images.SecureBaseUrl
}

func (tmdbConfiguration *TheMovieDbConfiguration) GetImageSizes(ctx context.Context, imageType string) []string {
    configuration := tmdbConfiguration.getConfiguration(ctx)

    switch imageType {
        case "backdrop":
            return configuration.Images.BackdropSizes
        case "logo":
            return configuration.Images.LogoSizes
        default:
            return configuration.Images.PosterSizes
    }
}

func (tmdbConfiguration *TheMovieDbConfiguration) getConfiguration(ctx context.Context) ApiConfigurationResult {
    var config ApiConfigurationResult

    exists, err := tmdbConfiguration.cache.Get(ctx, CacheKey, &config)
    if err != nil {
        panic(err)
    }

    if exists {
        return config
    }

    config, err = tmdbConfiguration.client.GetConfiguration()
    if err != nil {
        panic(err)
    }

    if err = tmdbConfiguration.cache.Set(ctx, CacheKey, config, CacheLifetime); err != nil {
        panic(err)
    }

    return config
}
