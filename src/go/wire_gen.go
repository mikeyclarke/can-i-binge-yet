// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
	cache2 "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/asset"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/cache"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/config"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/file"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/front_end/app/controllers"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/id_generator"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/show"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/template"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
	"github.com/mikeyclarke/can-i-binge-yet/src/go/url"
	"github.com/nikolalohinski/gonja"
	config2 "github.com/nikolalohinski/gonja/config"
	"github.com/nikolalohinski/gonja/loaders"
	"time"
)

// Injectors from container.go:

func CreateHomeController() *controllers.HomeController {
	applicationConfig := config.NewApplicationConfig()
	redisHost := applicationConfig.RedisHost
	redisPort := applicationConfig.RedisPort
	redisPassword := applicationConfig.RedisPassword
	cacheCache := NewRedisCache(redisHost, redisPort, redisPassword)
	tmdbApiKey := applicationConfig.TmdbApiKey
	tmdbApiBaseUrl := applicationConfig.TmdbApiBaseUrl
	theMovieDbClient := themoviedb.NewTheMovieDbClient(tmdbApiKey, tmdbApiBaseUrl)
	theMovieDbConfiguration := themoviedb.NewTheMovieDbConfiguration(cacheCache, theMovieDbClient)
	showImageFormatter := show.NewShowImageFormatter(theMovieDbConfiguration)
	slugGenerator := url.NewSlugGenerator()
	showSearch := show.NewShowSearch(showImageFormatter, slugGenerator, theMovieDbClient)
	cache2 := cache.NewCache(cacheCache)
	trendingShows := show.NewTrendingShows(cache2, showImageFormatter, slugGenerator, theMovieDbClient)
	homeController := controllers.NewHomeController(showSearch, trendingShows)
	return homeController
}

func CreateTemplateRenderer() *template.TemplateRenderer {
	applicationConfig := config.NewApplicationConfig()
	assetDirectory := applicationConfig.AssetDirectory
	assetManifestPath := applicationConfig.AssetManifestPath
	assetManifest := asset.NewAssetManifest(assetManifestPath)
	fileReader := file.NewFileReader()
	alphanumericIdGenerator := id_generator.NewAlphanumericIdGenerator()
	assetRenderer := asset.NewAssetRenderer(assetDirectory, assetManifest, fileReader, alphanumericIdGenerator)
	templatesDirectory := applicationConfig.TemplatesDirectory
	environment := NewGonjaEnvironment(templatesDirectory)
	templateRenderer := template.NewTemplateRenderer(applicationConfig, assetRenderer, environment, templatesDirectory)
	return templateRenderer
}

// container.go:

func NewGonjaEnvironment(rootDir config.TemplatesDirectory) *gonja.Environment {
	loader := loaders.MustNewFileSystemLoader(string(rootDir))
	environment := gonja.NewEnvironment(config2.DefaultConfig, loader)
	evalLoader := template.NewGonjaCachedEvalTemplateLoader(environment)
	environment.EvalConfig.Loader = evalLoader
	return environment
}

func NewRedisCache(host config.RedisHost, port config.RedisPort, pw config.RedisPassword) *cache2.Cache {
	redis2 := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", string(host), int(port)),
		Password: string(pw),
	})
	cache3 := cache2.New(&cache2.Options{
		Redis:      redis2,
		LocalCache: cache2.NewTinyLFU(1000, time.Minute),
	})
	return cache3
}
