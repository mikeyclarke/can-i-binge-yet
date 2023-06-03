package show

import (
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type SeasonEpisode struct {
    Name string
    Number int
    AirDate *time.Time
    HasAired bool
}

type SeasonEpisodesFormatterInterface interface {
    Format(episodes []themoviedb.ApiSeasonEpisodeResult) []SeasonEpisode
}
