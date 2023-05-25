package themoviedb

import (
    "context"
    "time"

    "github.com/go-redis/cache/v8"
)

const (
    CacheKey = "tmdb_configuration"
    CacheLifetime = time.Hour * 120 // 5 days
)

type TheMovieDbConfiguration struct {
    cache *cache.Cache
    client *TheMovieDbClient
}

func NewTheMovieDbConfiguration(cache *cache.Cache, tmdbClient *TheMovieDbClient) *TheMovieDbConfiguration {
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

    err := tmdbConfiguration.cache.Get(ctx, CacheKey, &config)
    if err == nil {
        return config
    }

    if err != cache.ErrCacheMiss {
        panic(err)
    }

    config = tmdbConfiguration.client.GetConfiguration()
    if err = tmdbConfiguration.cache.Set(&cache.Item{
        Ctx: ctx,
        Key: CacheKey,
        Value: config,
        TTL: CacheLifetime,
    }); err != nil {
        panic(err)
    }

    return config
}
