package show

type SeasonEpisodesLoaderInterface interface {
    Load(showId int, seasonNumber int) *[]SeasonEpisode
}
