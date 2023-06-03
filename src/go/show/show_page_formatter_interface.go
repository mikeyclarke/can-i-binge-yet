package show

import (
    "context"
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type Network struct {
    Logo string
    Name string
}

type Season struct {
    Name string
    AirDate *time.Time
    Year int
    EpisodeCount int
    Episodes []SeasonEpisode
    SeasonNumber int
    HasStarted bool
}

type ShowPageResult struct {
    TmdbId int
    Title string
    Tagline string
    CountriesEmoji []string
    Overview string
    PosterImage *ShowImageResult
    BackdropImage *ShowImageResult
    Slug string
    UrlPath string
    StatusText string
    Networks []Network
    Seasons []Season
    AirDatesDisplay string
    HasEnded bool
    IsBingeable bool
}

type ShowPageFormatterInterface interface {
    Format(ctx context.Context, show themoviedb.ApiShowResult, lastSeason *themoviedb.ApiSeasonResult) *ShowPageResult
}
