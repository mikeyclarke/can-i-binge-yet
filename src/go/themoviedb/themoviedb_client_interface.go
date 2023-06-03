package themoviedb

import (
    "fmt"
)

type ApiConfigurationResult struct {
    Images struct {
        BaseUrl string `json:"base_url"`
        SecureBaseUrl string `json:"secure_base_url"`
        BackdropSizes []string `json:"backdrop_sizes"`
        LogoSizes []string `json:"logo_sizes"`
        PosterSizes []string `json:"poster_sizes"`
    } `json:"images"`
}

type ApiTrendingShowResult struct {
    Adult bool `json:"adult"`
    Id int `json:"id"`
    Name string `json:"name"`
    PosterPath *string `json:"poster_path"`
}

type ApiTrendingShowsResult struct {
    Results []ApiTrendingShowResult `json:"results"`
}

type ApiSearchShowResult struct {
    Id int `json:"id"`
    FirstAirDate string `json:"first_air_date"`
    Name string `json:"name"`
    OriginCountry []string `json:"origin_country"`
    Overview string `json:"overview"`
    PosterPath *string `json:"poster_path"`
}

type ApiSearchShowsResult struct {
    Page int `json:"page"`
    Results []ApiSearchShowResult `json:"results"`
    TotalPages int `json:"total_pages"`
    TotalResults int `json:"total_results"`
}

type ApiSeasonEpisodeResult struct {
    AirDate string `json:"air_date"`
    EpisodeNumber int `json:"episode_number"`
    Name string `json:"name"`
}

type ApiSeasonResult struct {
    AirDate string `json:"air_date"`
    EpisodeCount int `json:"episode_count"`
    Episodes []ApiSeasonEpisodeResult `json:"episodes"`
    Name string `json:"name"`
    SeasonNumber int `json:"season_number"`
}

type ApiNetworkResult struct {
    Name string `json:"name"`
    LogoPath string `json:"logo_path"`
}

type ApiShowResultEmbededSeason struct {
    AirDate string `json:"air_date"`
    EpisodeCount int `json:"episode_count"`
    Name string `json:"name"`
    SeasonNumber int `json:"season_number"`
}

type ApiShowResult struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Tagline string `json:"tagline"`
    OriginCountry []string `json:"origin_country"`
    Overview string `json:"overview"`
    Networks []ApiNetworkResult `json:"networks"`
    PosterPath *string `json:"poster_path"`
    BackdropPath *string `json:"backdrop_path"`
    Status string `json:"status"`
    FirstAirDate string `json:"first_air_date"`
    Seasons []ApiShowResultEmbededSeason `json:"seasons"`
}

type TheMovieDbClientResponseError struct {
    status string
    statusCode int
}

func (err *TheMovieDbClientResponseError) Error() string {
    return fmt.Sprintf("Request failed with status: %s", err.status)
}

func (err *TheMovieDbClientResponseError) StatusCode() int {
    return err.statusCode
}

type TheMovieDbClientInterface interface {
    GetConfiguration() (ApiConfigurationResult, error)
    GetTrendingShows(timeWindow string) (ApiTrendingShowsResult, error)
    SearchShows(searchToken string, page int) (ApiSearchShowsResult, error)
    GetShow(id int) (ApiShowResult, error)
    GetShowSeason(id int, seasonNumber int) (ApiSeasonResult, error)
}
