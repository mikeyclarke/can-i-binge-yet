package themoviedb

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

type TheMovieDbClientInterface interface {
    GetConfiguration() ApiConfigurationResult
    GetTrendingShows(timeWindow string) ApiTrendingShowsResult
    SearchShows(searchToken string, page int) ApiSearchShowsResult
}
