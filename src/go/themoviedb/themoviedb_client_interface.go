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
    Id int `json:"id"`
    Name string `json:"name"`
    PosterPath *string `json:"poster_path"`
}

type ApiTrendingShowsResult struct {
    Results []ApiTrendingShowResult `json:"results"`
}

type TheMovieDbClientInterface interface {
    GetConfiguration() ApiConfigurationResult
    GetTrendingShows(timeWindow string) ApiTrendingShowsResult
}
