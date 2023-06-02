package show

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/cache"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/mocks/context"
)

func TestGetAllWithPopulatedCache(t *testing.T) {
    cache := cache.NewMockCacheInterface(t)
    showImageFormatter := NewMockShowImageFormatterInterface(t)
    slugGenerator := url.NewMockSlugGeneratorInterface(t)
    tmdbClient := themoviedb.NewMockTheMovieDbClientInterface(t)
    trendingShows := NewTrendingShows(cache, showImageFormatter, slugGenerator, tmdbClient)

    sandmanPosterPath := "/q54qEgagGOYCq5D1903eBVMNkbo.jpg"
    strangerThingsPosterPath := "/49WJfeN0moxb9IPfGn8AIqMGskD.jpg"
    adultShowPosterPath := "/tg7jksahekowl.jpg"
    apiResults := []themoviedb.ApiTrendingShowResult{
        themoviedb.ApiTrendingShowResult{
            Adult: false,
            Id: 90802,
            Name: "The Sandman",
            PosterPath: &sandmanPosterPath,
        },
        themoviedb.ApiTrendingShowResult{
            Adult: false,
            Id: 66732,
            Name: "Stranger Things",
            PosterPath: &strangerThingsPosterPath,
        },
        themoviedb.ApiTrendingShowResult{
            Adult: true,
            Id: 78434,
            Name: "Some Adult Show",
            PosterPath: &adultShowPosterPath,
        },
    }

    ctx := context.NewMockContext(t)
    sandmanPoster := ShowImageResult{
        "sandman-default-poster.jpg",
        "sandman-srcsets.jpg 92w",
    }
    strangerThingsPoster := ShowImageResult{
        "stranger-things-default-poster.jpg",
        "stranger-things-srcsets.jpg 92w",
    }

    cache.On("Get", ctx, CacheKey, mock.AnythingOfType("*themoviedb.ApiTrendingShowsResult")).
        Return(true, nil).Run(func(args mock.Arguments) {
            arg := args.Get(2).(*themoviedb.ApiTrendingShowsResult)
            arg.Results = apiResults
        })
    slugGenerator.On("Generate", apiResults[0].Name, 0).Return("the-sandman")
    slugGenerator.On("Generate", apiResults[1].Name, 0).Return("stranger-things")
    showImageFormatter.On("Format", ctx, "poster", *apiResults[0].PosterPath, "").Return(sandmanPoster)
    showImageFormatter.On("Format", ctx, "poster", *apiResults[1].PosterPath, "").Return(strangerThingsPoster)

    expected := []TrendingShowResult{
        TrendingShowResult{
            TmdbId: apiResults[0].Id,
            Title: apiResults[0].Name,
            Slug: "the-sandman",
            UrlPath: fmt.Sprintf("%d-%s", apiResults[0].Id, "the-sandman"),
            PosterImage: sandmanPoster,
        },
        TrendingShowResult{
            TmdbId: apiResults[1].Id,
            Title: apiResults[1].Name,
            Slug: "stranger-things",
            UrlPath: fmt.Sprintf("%d-%s", apiResults[1].Id, "stranger-things"),
            PosterImage: strangerThingsPoster,
        },
    }

    result := trendingShows.GetAll(ctx)
    assert.Equal(t, expected, result)
}

func TestGetAllWithEmptyCache(t *testing.T) {
    cache := cache.NewMockCacheInterface(t)
    showImageFormatter := NewMockShowImageFormatterInterface(t)
    slugGenerator := url.NewMockSlugGeneratorInterface(t)
    tmdbClient := themoviedb.NewMockTheMovieDbClientInterface(t)
    trendingShows := NewTrendingShows(cache, showImageFormatter, slugGenerator, tmdbClient)

    sandmanPosterPath := "/q54qEgagGOYCq5D1903eBVMNkbo.jpg"
    strangerThingsPosterPath := "/49WJfeN0moxb9IPfGn8AIqMGskD.jpg"
    adultShowPosterPath := "/tg7jksahekowl.jpg"
    apiResults := []themoviedb.ApiTrendingShowResult{
        themoviedb.ApiTrendingShowResult{
            Adult: false,
            Id: 90802,
            Name: "The Sandman",
            PosterPath: &sandmanPosterPath,
        },
        themoviedb.ApiTrendingShowResult{
            Adult: false,
            Id: 66732,
            Name: "Stranger Things",
            PosterPath: &strangerThingsPosterPath,
        },
        themoviedb.ApiTrendingShowResult{
            Adult: true,
            Id: 78434,
            Name: "Some Adult Show",
            PosterPath: &adultShowPosterPath,
        },
    }
    apiResult := themoviedb.ApiTrendingShowsResult{
        Results: apiResults,
    }

    ctx := context.NewMockContext(t)
    sandmanPoster := ShowImageResult{
        "sandman-default-poster.jpg",
        "sandman-srcsets.jpg 92w",
    }
    strangerThingsPoster := ShowImageResult{
        "stranger-things-default-poster.jpg",
        "stranger-things-srcsets.jpg 92w",
    }

    cache.On("Get", ctx, CacheKey, mock.AnythingOfType("*themoviedb.ApiTrendingShowsResult")).Return(false, nil)
    tmdbClient.On("GetTrendingShows", "").Return(apiResult)
    cache.On("Set", ctx, CacheKey, apiResult, CacheLifetime).Return(nil)
    slugGenerator.On("Generate", apiResults[0].Name, 0).Return("the-sandman")
    slugGenerator.On("Generate", apiResults[1].Name, 0).Return("stranger-things")
    showImageFormatter.On("Format", ctx, "poster", *apiResults[0].PosterPath, "").Return(sandmanPoster)
    showImageFormatter.On("Format", ctx, "poster", *apiResults[1].PosterPath, "").Return(strangerThingsPoster)

    expected := []TrendingShowResult{
        TrendingShowResult{
            TmdbId: apiResults[0].Id,
            Title: apiResults[0].Name,
            Slug: "the-sandman",
            UrlPath: fmt.Sprintf("%d-%s", apiResults[0].Id, "the-sandman"),
            PosterImage: sandmanPoster,
        },
        TrendingShowResult{
            TmdbId: apiResults[1].Id,
            Title: apiResults[1].Name,
            Slug: "stranger-things",
            UrlPath: fmt.Sprintf("%d-%s", apiResults[1].Id, "stranger-things"),
            PosterImage: strangerThingsPoster,
        },
    }

    result := trendingShows.GetAll(ctx)
    assert.Equal(t, expected, result)
}
