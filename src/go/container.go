//+build wireinject

package main

import (
    "fmt"
    "time"

    "github.com/google/wire"
    "github.com/nikolalohinski/gonja"
    gonjaConfig "github.com/nikolalohinski/gonja/config"
    gonjaLoaders "github.com/nikolalohinski/gonja/loaders"
    "github.com/go-redis/redis/v8"
    redisCache "github.com/go-redis/cache/v8"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/asset"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/cache"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/file"
    idGenerator "github.com/mikeyclarke/can-i-binge-yet/src/go/id_generator"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/template"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/front_end/app/controllers"
)

func NewGonjaEnvironment(rootDir config.TemplatesDirectory) *gonja.Environment {
    loader := gonjaLoaders.MustNewFileSystemLoader(string(rootDir))
    environment := gonja.NewEnvironment(gonjaConfig.DefaultConfig, loader)
    evalLoader := template.NewGonjaCachedEvalTemplateLoader(environment)
    environment.EvalConfig.Loader = evalLoader
    return environment
}

func NewRedisCache(host config.RedisHost, port config.RedisPort, pw config.RedisPassword) *redisCache.Cache {
    redis := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", string(host), int(port)),
        Password: string(pw),
    })
    cache := redisCache.New(&redisCache.Options{
        Redis: redis,
        LocalCache: redisCache.NewTinyLFU(1000, time.Minute),
    })
    return cache
}

func CreateHomeController() *controllers.HomeController {
    panic(wire.Build(
        config.NewApplicationConfig,
        wire.FieldsOf(
            new(config.ApplicationConfig),
            "RedisHost",
            "RedisPort",
            "RedisPassword",
            "TmdbApiBaseUrl",
            "TmdbApiKey",
        ),
        NewRedisCache,
        wire.Bind(new(cache.CacheInterface), new(*cache.Cache)),
        cache.NewCache,
        wire.Bind(new(themoviedb.TheMovieDbConfigurationInterface), new(*themoviedb.TheMovieDbConfiguration)),
        wire.Bind(new(themoviedb.TheMovieDbClientInterface), new(*themoviedb.TheMovieDbClient)),
        themoviedb.NewTheMovieDbClient,
        themoviedb.NewTheMovieDbConfiguration,
        wire.Bind(new(url.SlugGeneratorInterface), new(*url.SlugGenerator)),
        url.NewSlugGenerator,
        wire.Bind(new(show.ShowImageFormatterInterface), new(*show.ShowImageFormatter)),
        show.NewShowImageFormatter,
        show.NewTrendingShows,
        show.NewShowSearch,
        controllers.NewHomeController,
    ))
}

func CreateTemplateRenderer() *template.TemplateRenderer {
    panic(wire.Build(
        config.NewApplicationConfig,
        wire.FieldsOf(
            new(config.ApplicationConfig),
            "AssetDirectory",
            "AssetManifestPath",
            "TemplatesDirectory",
        ),
        NewGonjaEnvironment,
        wire.Bind(new(asset.AssetManifestInterface), new(*asset.AssetManifest)),
        asset.NewAssetManifest,
        asset.NewAssetRenderer,
        wire.Bind(new(file.FileReaderInterface), new(*file.FileReader)),
        file.NewFileReader,
        wire.Bind(new(idGenerator.AlphanumericIdGeneratorInterface), new(*idGenerator.AlphanumericIdGenerator)),
        idGenerator.NewAlphanumericIdGenerator,
        template.NewTemplateRenderer,
    ))
}
