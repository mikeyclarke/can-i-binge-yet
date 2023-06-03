package show

import (
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type SeasonEpisodesFormatter struct {}

func NewSeasonEpisodesFormatter() *SeasonEpisodesFormatter {
    return &SeasonEpisodesFormatter{}
}

func (formatter *SeasonEpisodesFormatter) Format(episodes []themoviedb.ApiSeasonEpisodeResult) []SeasonEpisode {
    today := time.Now()
    var formatted []SeasonEpisode

    for _, episode := range episodes {
        formattedEpisode := SeasonEpisode{
            Name: episode.Name,
            Number: episode.EpisodeNumber,
        }

        var hasAired bool
        if episode.AirDate != "" {
            airDate, err := time.Parse(time.DateOnly, episode.AirDate)
            if err != nil {
                panic(err)
            }
            formattedEpisode.AirDate = &airDate
            hasAired = (today.Compare(airDate) != -1)
        }
        formattedEpisode.HasAired = hasAired

        formatted = append(formatted, formattedEpisode)
    }

    return formatted
}
