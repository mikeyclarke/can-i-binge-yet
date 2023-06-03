package show

import (
    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type SeasonEpisodesLoader struct {
    seasonEpisodesFormatter SeasonEpisodesFormatterInterface
    tmdbClient themoviedb.TheMovieDbClientInterface
}

func NewSeasonEpisodesLoader(
    seasonEpisodesFormatter SeasonEpisodesFormatterInterface,
    tmdbClient themoviedb.TheMovieDbClientInterface,
) *SeasonEpisodesLoader {
    return &SeasonEpisodesLoader{seasonEpisodesFormatter, tmdbClient}
}

func (loader *SeasonEpisodesLoader) Load(showId int, seasonNumber int) *[]SeasonEpisode {
    season, err := loader.tmdbClient.GetShowSeason(showId, seasonNumber)
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

    episodes := loader.seasonEpisodesFormatter.Format(season.Episodes)
    return &episodes
}
