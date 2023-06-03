package show

import (
    "context"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type ShowLoader struct {
    showPageFormatter ShowPageFormatterInterface
    tmdbClient themoviedb.TheMovieDbClientInterface
}

func NewShowLoader(
    showPageFormatter ShowPageFormatterInterface,
    tmdbClient themoviedb.TheMovieDbClientInterface,
) *ShowLoader {
    return &ShowLoader{showPageFormatter, tmdbClient}
}

func (loader *ShowLoader) Load(ctx context.Context, id int) *ShowPageResult {
    show, err := loader.tmdbClient.GetShow(id)
    if err != nil {
        switch err {
            case err.(*themoviedb.TheMovieDbClientResponseError):
                responseErr := err.(*themoviedb.TheMovieDbClientResponseError)
                if responseErr.StatusCode() == 404 {
                    return nil
                }
                panic(err)
            default:
                panic(err)
        }
    }

    show.Seasons = loader.filterSeasons(show)

    var lastSeason *themoviedb.ApiSeasonResult
    if len(show.Seasons) != 0 {
        seasonNumber := show.Seasons[len(show.Seasons) - 1].SeasonNumber
        season, err := loader.tmdbClient.GetShowSeason(id, seasonNumber)
        if err != nil {
            panic(err)
        }
        lastSeason = &season
    }

    return loader.showPageFormatter.Format(ctx, show, lastSeason)
}

func (loader *ShowLoader) filterSeasons(show themoviedb.ApiShowResult) []themoviedb.ApiShowResultEmbededSeason {
    var filtered []themoviedb.ApiShowResultEmbededSeason

    for _, season := range show.Seasons {
        if season.EpisodeCount > 0 {
            filtered = append(filtered, season)
        }
    }

    return filtered
}
