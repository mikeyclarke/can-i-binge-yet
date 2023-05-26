package themoviedb

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "time"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/config"
)

type TheMovieDbClient struct {
    apiKey config.TmdbApiKey
    baseUrl config.TmdbApiBaseUrl
    httpClient *http.Client
}

func NewTheMovieDbClient(apiKey config.TmdbApiKey, baseUrl config.TmdbApiBaseUrl) *TheMovieDbClient {
    httpClient := &http.Client{
        Timeout: time.Second * 15,
    }
    return &TheMovieDbClient{apiKey, baseUrl, httpClient}
}

func (tmdbClient *TheMovieDbClient) GetConfiguration() ApiConfigurationResult {
    var result ApiConfigurationResult
    tmdbClient.makeRequest("/3/configuration", map[string]string{}, &result)
    return result
}

func (tmdbClient *TheMovieDbClient) GetTrendingShows(timeWindow string) ApiTrendingShowsResult {
    if timeWindow == "" {
        timeWindow = "day"
    }
    var result ApiTrendingShowsResult
    tmdbClient.makeRequest(fmt.Sprintf("/3/trending/tv/%s", timeWindow), map[string]string{}, &result)
    return result
}

func (tmdbClient *TheMovieDbClient) SearchShows(searchToken string, page int) ApiSearchShowsResult {
    if page == 0 {
        page = 1
    }

    var result ApiSearchShowsResult
    tmdbClient.makeRequest("/3/search/tv", map[string]string{"query": searchToken}, &result)
    return result
}

func (tmdbClient *TheMovieDbClient) makeRequest(endpoint string, params map[string]string, result any) {
    requestParams := url.Values{}
    for key, val := range params {
        requestParams.Add(key, val)
    }

    requestParams.Set("api_key", string(tmdbClient.apiKey))

    requestUrl := fmt.Sprintf("%s%s?%s", string(tmdbClient.baseUrl), endpoint, requestParams.Encode())
    response, err := tmdbClient.httpClient.Get(requestUrl)
    if err != nil {
        panic(err)
    }

    body, err := io.ReadAll(response.Body)
    response.Body.Close()

    if response.StatusCode < 200 || response.StatusCode > 299 || err != nil {
        panic(err)
    }

    err = json.Unmarshal(body, &result)
    if err != nil {
        panic(err)
    }
}
