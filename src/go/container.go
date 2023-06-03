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
    apiAppControllers "github.com/mikeyclarke/can-i-binge-yet/src/go/api/app/controllers"
    frontEndAppControllers "github.com/mikeyclarke/can-i-binge-yet/src/go/front_end/app/controllers"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/front_end/app/middleware"
)

func NewGonjaEnvironment(rootDir config.TemplatesDirectory) *gonja.Environment {
    loader := gonjaLoaders.MustNewFileSystemLoader(string(rootDir))
    environment := gonja.NewEnvironment(gonjaConfig.DefaultConfig, loader)
    evalLoader := template.NewGonjaCachedEvalTemplateLoader(environment)
    environment.EvalConfig.Loader = evalLoader
    for name, filterFunc := range template.Filters {
        environment.Filters.Register(name, filterFunc)
    }
    environment.Globals.Merge(template.Functions)
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

var ConfigSet = wire.NewSet(
    config.NewApplicationConfig,
    wire.FieldsOf(
        new(config.ApplicationConfig),
        "AssetDirectory",
        "AssetManifestPath",
        "RedisHost",
        "RedisPort",
        "RedisPassword",
        "TemplatesDirectory",
        "TmdbApiBaseUrl",
        "TmdbApiKey",
    ),
)

var CacheSet = wire.NewSet(
    NewRedisCache,
    wire.Bind(new(cache.CacheInterface), new(*cache.Cache)),
    cache.NewCache,
)

var TemplateRendererSet = wire.NewSet(
    NewGonjaEnvironment,
    wire.Bind(new(asset.AssetManifestInterface), new(*asset.AssetManifest)),
    asset.NewAssetManifest,
    asset.NewAssetRenderer,
    wire.Bind(new(file.FileReaderInterface), new(*file.FileReader)),
    file.NewFileReader,
    wire.Bind(new(idGenerator.AlphanumericIdGeneratorInterface), new(*idGenerator.AlphanumericIdGenerator)),
    idGenerator.NewAlphanumericIdGenerator,
    template.NewTemplateRenderer,
)

func CreateHomeController() *frontEndAppControllers.HomeController {
    panic(wire.Build(
        ConfigSet,
        CacheSet,
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
        frontEndAppControllers.NewHomeController,
    ))
}

func CreateShowController() *frontEndAppControllers.ShowController {
    panic(wire.Build(
        frontEndAppControllers.NewShowController,
    ))
}

func CreateSeasonEpisodesController() *apiAppControllers.SeasonEpisodesController {
    panic(wire.Build(
        ConfigSet,
        TemplateRendererSet,
        wire.Bind(new(themoviedb.TheMovieDbClientInterface), new(*themoviedb.TheMovieDbClient)),
        themoviedb.NewTheMovieDbClient,
        wire.Bind(new(show.SeasonEpisodesLoaderInterface), new(*show.SeasonEpisodesLoader)),
        show.NewSeasonEpisodesLoader,
        wire.Bind(new(show.SeasonEpisodesFormatterInterface), new(*show.SeasonEpisodesFormatter)),
        show.NewSeasonEpisodesFormatter,
        apiAppControllers.NewSeasonEpisodesController,
    ))
}

func CreateShowLoaderMiddleware() *middleware.ShowLoaderMiddleware {
    panic(wire.Build(
        ConfigSet,
        CacheSet,
        wire.Bind(new(themoviedb.TheMovieDbConfigurationInterface), new(*themoviedb.TheMovieDbConfiguration)),
        wire.Bind(new(themoviedb.TheMovieDbClientInterface), new(*themoviedb.TheMovieDbClient)),
        themoviedb.NewTheMovieDbClient,
        themoviedb.NewTheMovieDbConfiguration,
        wire.Bind(new(url.SlugGeneratorInterface), new(*url.SlugGenerator)),
        url.NewSlugGenerator,
        wire.Bind(new(show.ShowImageFormatterInterface), new(*show.ShowImageFormatter)),
        show.NewShowImageFormatter,
        wire.Bind(new(show.SeasonEpisodesFormatterInterface), new(*show.SeasonEpisodesFormatter)),
        show.NewSeasonEpisodesFormatter,
        wire.Bind(new(show.ShowPageFormatterInterface), new(*show.ShowPageFormatter)),
        show.NewShowPageFormatter,
        wire.Bind(new(show.ShowLoaderInterface), new(*show.ShowLoader)),
        show.NewShowLoader,
        middleware.NewShowLoaderMiddleware,
    ))
}

func CreateTemplateRenderer() *template.TemplateRenderer {
    panic(wire.Build(
        ConfigSet,
        TemplateRendererSet,
    ))
}
