package show

import (
    "context"
    "fmt"
    "strconv"
    "time"

    "github.com/enescakir/emoji"
    "golang.org/x/exp/slices"
    "golang.org/x/text/language"
    "golang.org/x/text/message"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
)

var ShowEndedStatuses = []string{"Cancelled", "Canceled", "Ended"}

type ShowPageFormatter struct {
    seasonEpisodesFormatter SeasonEpisodesFormatterInterface
    showImageFormatter ShowImageFormatterInterface
    slugGenerator url.SlugGeneratorInterface
}

func NewShowPageFormatter(
    seasonEpisodesFormatter SeasonEpisodesFormatterInterface,
    showImageFormatter ShowImageFormatterInterface,
    slugGenerator url.SlugGeneratorInterface,
) *ShowPageFormatter {
    return &ShowPageFormatter{seasonEpisodesFormatter, showImageFormatter, slugGenerator}
}

func (showPageFormatter *ShowPageFormatter) Format(
    ctx context.Context,
    show themoviedb.ApiShowResult,
    lastSeason *themoviedb.ApiSeasonResult,
) *ShowPageResult {
    id := show.Id
    slug := showPageFormatter.slugGenerator.Generate(show.Name, 0)
    urlPath := fmt.Sprintf("%d-%s", id, slug)

    var countriesEmoji []string
    for _, code := range show.OriginCountry {
        flag, err := emoji.CountryFlag(code)
        if err != nil {
            panic(err)
        }
        countriesEmoji = append(countriesEmoji, string(flag))
    }

    networks := showPageFormatter.formatNetworks(ctx, show)
    seasons := showPageFormatter.formatSeasons(show, lastSeason)

    hasEnded := slices.Contains(ShowEndedStatuses, show.Status)

    var airDatesDisplay string
    var bingeable bool
    var statusText string
    if lastSeason != nil && len(lastSeason.Episodes) > 0 {
        formattedLastSeason := seasons[0]
        lastEpisode := formattedLastSeason.Episodes[len(formattedLastSeason.Episodes) - 1]
        airDatesDisplay = showPageFormatter.formatAirDates(show, &lastEpisode, hasEnded)
        statusText = showPageFormatter.formatStatusText(&formattedLastSeason, &lastEpisode)
        bingeable = lastEpisode.HasAired
    }

    formatted := &ShowPageResult{
        TmdbId: id,
        Title: show.Name,
        Tagline: show.Tagline,
        Slug: slug,
        CountriesEmoji: countriesEmoji,
        Overview: show.Overview,
        UrlPath: urlPath,
        StatusText: statusText,
        Networks: networks,
        Seasons: seasons,
        AirDatesDisplay: airDatesDisplay,
        HasEnded: hasEnded,
        IsBingeable: bingeable,
    }

    if show.PosterPath != nil && *show.PosterPath != "" {
        poster := showPageFormatter.showImageFormatter.Format(ctx, "poster", *show.PosterPath, "")
        formatted.PosterImage = &poster
    }

    if show.BackdropPath != nil && *show.BackdropPath != "" {
        backdrop := showPageFormatter.showImageFormatter.Format(ctx, "backdrop", *show.BackdropPath, "")
        formatted.BackdropImage = &backdrop
    }

    return formatted
}

func (showPageFormatter *ShowPageFormatter) formatNetworks(
    ctx context.Context,
    show themoviedb.ApiShowResult,
) []Network {
    var networks []Network
    for _, apiNetwork := range show.Networks {
        var logo string
        if apiNetwork.LogoPath != "" {
            logo = showPageFormatter.showImageFormatter.FormatUrl(ctx, apiNetwork.LogoPath, "h60")
        }
        network := Network{
            Name: apiNetwork.Name,
            Logo: logo,
        }
        networks = append(networks, network)
    }
    return networks
}

func (showPageFormatter *ShowPageFormatter) formatSeasons(
    show themoviedb.ApiShowResult,
    lastSeason *themoviedb.ApiSeasonResult,
) []Season {
    today := time.Now()
    var seasons []Season

    for i := len(show.Seasons) -1; i >= 0; i-- {
        apiSeason := show.Seasons[i]
        season := Season{
            Name: apiSeason.Name,
            EpisodeCount: apiSeason.EpisodeCount,
            SeasonNumber: apiSeason.SeasonNumber,
        }

        if apiSeason.AirDate != "" {
            airDate, err := time.Parse(time.DateOnly, apiSeason.AirDate)
            if err != nil {
                panic(err)
            }

            season.AirDate = &airDate
            season.Year = airDate.Year()
            season.HasStarted = (today.Compare(airDate) != -1)
        }

        if lastSeason != nil && lastSeason.SeasonNumber == season.SeasonNumber {
            season.Episodes = showPageFormatter.seasonEpisodesFormatter.Format(lastSeason.Episodes)
        }

        seasons = append(seasons, season)
    }
    return seasons
}

func (showPageFormatter *ShowPageFormatter) formatAirDates(
    show themoviedb.ApiShowResult,
    lastEpisode *SeasonEpisode,
    hasEnded bool,
) string {
    var startYear string
    if show.FirstAirDate != "" {
        startDate, err := time.Parse(time.DateOnly, show.FirstAirDate)
        if err != nil {
            panic(err)
        }
        startYear = strconv.Itoa(startDate.Year())
    }

    if startYear == "" {
        return startYear
    }

    if !hasEnded {
        return fmt.Sprintf("%s - present", startYear)
    }

    if lastEpisode == nil || lastEpisode.AirDate == nil {
        return fmt.Sprintf("%s -", startYear)
    }

    endYear := strconv.Itoa(lastEpisode.AirDate.Year())

    if startYear == endYear {
        return startYear
    }

    return fmt.Sprintf("%s - %s", startYear, endYear)
}

func (showPageFormatter *ShowPageFormatter) formatStatusText(
    lastSeason *Season,
    lastEpisode *SeasonEpisode,
) string {
    if lastSeason == nil || lastEpisode == nil || lastSeason.AirDate == nil {
        return "We’re unable to obtain the air dates of this show right now"
    }

    lang := language.AmericanEnglish
    p := message.NewPrinter(lang)

    seasonNumber := lastSeason.SeasonNumber
    premiereDate := showPageFormatter.formatFriendlyDate(p, lastSeason.AirDate)

    var lastEpisodeDate string
    if lastEpisode.AirDate != nil {
        lastEpisodeDate = showPageFormatter.formatFriendlyDate(p, lastEpisode.AirDate)
    }

    var unairedEpisodes int
    for _, episode := range lastSeason.Episodes {
        if !episode.HasAired {
            unairedEpisodes += 1
        }
    }

    if lastEpisode.HasAired {
        return fmt.Sprintf("Season %d has concluded and all episodes have aired", seasonNumber)
    }

    if lastSeason.HasStarted {
        episodesLabel := "episodes"
        if unairedEpisodes == 1 {
            episodesLabel = "episode"
        }

        if lastEpisodeDate == "" {
            return fmt.Sprintf(
                "Season %d has %d %s left but its conclusion date is currently unknown",
                seasonNumber,
                unairedEpisodes,
                episodesLabel,
            )
        }

        return fmt.Sprintf(
            "Season %d has %d %s left and concludes %s",
            seasonNumber,
            unairedEpisodes,
            episodesLabel,
            lastEpisodeDate,
        )
    }

    if lastEpisodeDate == "" {
        return fmt.Sprintf(
            "Season %d premieres %s but its conclusion date is currently unknown",
            seasonNumber,
            premiereDate,
        )
    }

    return fmt.Sprintf("Season %d premieres %s and concludes %s", seasonNumber, premiereDate, lastEpisodeDate)
}

func (showPageFormatter *ShowPageFormatter) formatFriendlyDate(p *message.Printer, airDate *time.Time) string {
    today := time.Now()
    weekFromNow := today.AddDate(0, 0, 7)

    var pattern string
    if airDate.Compare(weekFromNow) == -1 {
        pattern = "|start_strong|%s|end_strong|"
    } else {
        pattern = "%s"
    }

    pattern += ", %s %d"

    result := p.Sprintf(pattern, airDate.Weekday(), airDate.Month(), airDate.Day())

    if airDate.Year() != today.Year() {
        result += p.Sprintf(" %s", strconv.Itoa(airDate.Year()))
    }

    return result
}
