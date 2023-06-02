package show

import (
    "context"
    "fmt"
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/cache"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
)

const (
    CacheKey = "tmdb_trending_shows"
    CacheLifetime = time.Hour * 24
)

type TrendingShowResult struct {
    TmdbId int
    Title string
    Slug string
    UrlPath string
    PosterImage ShowImageResult
}

type TrendingShows struct {
    cache cache.CacheInterface
    showImageFormatter ShowImageFormatterInterface
    slugGenerator url.SlugGeneratorInterface
    tmdbClient themoviedb.TheMovieDbClientInterface
}

func NewTrendingShows(
    cache cache.CacheInterface,
    showImageFormatter ShowImageFormatterInterface,
    slugGenerator url.SlugGeneratorInterface,
    tmdbClient themoviedb.TheMovieDbClientInterface,
) *TrendingShows {
    return &TrendingShows{cache, showImageFormatter, slugGenerator, tmdbClient}
}

func (trendingShows *TrendingShows) GetAll(ctx context.Context) []TrendingShowResult {
    shows := trendingShows.getResults(ctx)
    return trendingShows.formatResults(ctx, shows)
}

func (trendingShows *TrendingShows) getResults(ctx context.Context) themoviedb.ApiTrendingShowsResult {
    var shows themoviedb.ApiTrendingShowsResult

    exists, err := trendingShows.cache.Get(ctx, CacheKey, &shows)
    if err != nil {
        panic(err)
    }

    if exists {
        return shows
    }

    shows = trendingShows.tmdbClient.GetTrendingShows("")
    if err = trendingShows.cache.Set(ctx, CacheKey, shows, CacheLifetime); err != nil {
        panic(err)
    }

    return shows
}

func (trendingShows *TrendingShows) formatResults(
    ctx context.Context,
    shows themoviedb.ApiTrendingShowsResult,
) []TrendingShowResult {
    formatted := []TrendingShowResult{}

    for _, result := range shows.Results {
        if result.Adult {
            continue
        }

        id := result.Id
        slug := trendingShows.slugGenerator.Generate(result.Name, 0)
        urlPath := fmt.Sprintf("%d-%s", id, slug)
        posterImage := trendingShows.showImageFormatter.Format(ctx, "poster", *result.PosterPath, "")

        show := TrendingShowResult{
            id,
            result.Name,
            slug,
            urlPath,
            posterImage,
        }

        formatted = append(formatted, show)
    }

    return formatted
}
