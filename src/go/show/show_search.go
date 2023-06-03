package show

import (
    "context"
    "fmt"
    "time"

    "github.com/enescakir/emoji"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
)

type ShowSearchResult struct {
    TmdbId int
    Title string
    CountriesEmoji []string
    Year int
    Overview string
    PosterImage *ShowImageResult
    Slug string
    UrlPath string
}

type ShowSearchResults struct {
    Page int
    Shows []ShowSearchResult
    TotalResults int
    TotalPages int
}

type ShowSearch struct {
    showImageFormatter ShowImageFormatterInterface
    slugGenerator url.SlugGeneratorInterface
    tmdbClient themoviedb.TheMovieDbClientInterface
}

func NewShowSearch(
    showImageFormatter ShowImageFormatterInterface,
    slugGenerator url.SlugGeneratorInterface,
    tmdbClient themoviedb.TheMovieDbClientInterface,
) *ShowSearch {
    return &ShowSearch{showImageFormatter, slugGenerator, tmdbClient}
}

func (showSearch *ShowSearch) Search(ctx context.Context, searchToken string, page int) ShowSearchResults {
    if page == 0 {
        page = 1
    }

    shows, err := showSearch.tmdbClient.SearchShows(searchToken, page)
    if err != nil {
        panic(err)
    }
    return showSearch.formatResults(ctx, shows)
}

func (showSearch *ShowSearch) formatResults(
    ctx context.Context,
    shows themoviedb.ApiSearchShowsResult,
) ShowSearchResults {
    formatted := ShowSearchResults{
        Page: shows.Page,
        Shows: []ShowSearchResult{},
        TotalResults: shows.TotalResults,
        TotalPages: shows.TotalPages,
    }

    for _, result := range shows.Results {
        id := result.Id
        slug := showSearch.slugGenerator.Generate(result.Name, 0)
        urlPath := fmt.Sprintf("%d-%s", id, slug)

        var year int
        if result.FirstAirDate != "" {
            airDate, err := time.Parse(time.DateOnly, result.FirstAirDate)
            if err != nil {
                panic(err)
            }
            year = airDate.Year()
        }

        var countriesEmoji []string
        for _, code := range result.OriginCountry {
            flag, err := emoji.CountryFlag(code)
            if err != nil {
                panic(err)
            }
            countriesEmoji = append(countriesEmoji, string(flag))
        }

        show := ShowSearchResult{
            TmdbId: id,
            Title: result.Name,
            CountriesEmoji: countriesEmoji,
            Year: year,
            Overview: result.Overview,
            Slug: slug,
            UrlPath: urlPath,
        }

        if result.PosterPath != nil && *result.PosterPath != "" {
            posterImage := showSearch.showImageFormatter.Format(ctx, "poster", *result.PosterPath, "")
            show.PosterImage = &posterImage
        }

        formatted.Shows = append(formatted.Shows, show)
    }

    return formatted
}
