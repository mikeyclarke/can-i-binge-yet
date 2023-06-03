package show

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/mocks/context"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/url"
)

type ShowSearchTestSuite struct {
    suite.Suite
    showImageFormatter *MockShowImageFormatterInterface
    slugGenerator *url.MockSlugGeneratorInterface
    tmdbClient *themoviedb.MockTheMovieDbClientInterface
    showSearch *ShowSearch
    ctx *context.MockContext
}

func (suite *ShowSearchTestSuite) SetupTest() {
    suite.showImageFormatter = NewMockShowImageFormatterInterface(suite.T())
    suite.slugGenerator = url.NewMockSlugGeneratorInterface(suite.T())
    suite.tmdbClient = themoviedb.NewMockTheMovieDbClientInterface(suite.T())
    suite.showSearch = NewShowSearch(
        suite.showImageFormatter,
        suite.slugGenerator,
        suite.tmdbClient,
    )
    suite.ctx = context.NewMockContext(suite.T())
}

func (suite *ShowSearchTestSuite) TestSearch() {
    searchToken := "loot"
    page := 0

    apiShow1PosterPath := "/jf9EyK7KgHIAzpOEVujAFScUkcW.jpg"
    apiShow3PosterPath := "/4XU5Jr49lXhz0sTdttfphs52WY5.jpg"
    apiShows := []themoviedb.ApiSearchShowResult{
        themoviedb.ApiSearchShowResult{
            Id: 197449,
            FirstAirDate: "2022-06-23",
            Name: "Loot",
            OriginCountry: []string{"US"},
            Overview: "After divorcing her husband of 20 years, Molly Novak must figure out what to do …",
            PosterPath: &apiShow1PosterPath,
        },
        themoviedb.ApiSearchShowResult{
            Id: 204039,
            FirstAirDate: "2022-05-15",
            Name: "Loot - Blood Treasure",
            OriginCountry: []string{"AU"},
            Overview: "This series documents how the trade in stolen antiquities and art has become a …",
        },
        themoviedb.ApiSearchShowResult{
            Id: 156200,
            FirstAirDate: "2018-12-25",
            Name: "Hard Looters",
            OriginCountry: []string{"FR"},
            Overview: "",
            PosterPath: &apiShow3PosterPath,
        },
        themoviedb.ApiSearchShowResult{
            Id: 38900,
            FirstAirDate: "",
            Name: "Looteri Dulhan",
            OriginCountry: []string{"IO"},
            Overview: "Looteri Dulhan was a unique fiction serial in Imagine TV. It is basically the …",
        },
    }
    apiResults := themoviedb.ApiSearchShowsResult{
        Page: 1,
        Results: apiShows,
        TotalPages: 1,
        TotalResults: 4,
    }

    show1Slug := "loot"
    show1PosterImage := ShowImageResult{
        Default: "loot.jpg",
        Srcset: "loot-small.jpg 72w, loot-med.jpg 142w",
    }
    show2Slug := "loot-blood-treasure"
    show3Slug := "hard-looters"
    show3PosterImage := ShowImageResult{
        Default: "hard-looters.jpg",
        Srcset: "hard-looters-small.jpg 72w, hard-looters-med.jpg 142w",
    }
    show4Slug := "looteri-dulhan"

    suite.tmdbClient.On("SearchShows", searchToken, 1).Return(apiResults, nil)
    suite.slugGenerator.On("Generate", apiShows[0].Name, 0).Return(show1Slug)
    suite.showImageFormatter.On("Format", suite.ctx, "poster", *apiShows[0].PosterPath, "").Return(show1PosterImage)
    suite.slugGenerator.On("Generate", apiShows[1].Name, 0).Return(show2Slug)
    suite.slugGenerator.On("Generate", apiShows[2].Name, 0).Return(show3Slug)
    suite.showImageFormatter.On("Format", suite.ctx, "poster", *apiShows[2].PosterPath, "").Return(show3PosterImage)
    suite.slugGenerator.On("Generate", apiShows[3].Name, 0).Return(show4Slug)

    expected := ShowSearchResults{
        Page: 1,
        Shows: []ShowSearchResult{
            ShowSearchResult{
                TmdbId: apiShows[0].Id,
                Title: apiShows[0].Name,
                CountriesEmoji: []string{"🇺🇸"},
                Year: 2022,
                Overview: apiShows[0].Overview,
                PosterImage: &show1PosterImage,
                Slug: show1Slug,
                UrlPath: fmt.Sprintf("%d-%s", apiShows[0].Id, show1Slug),
            },
            ShowSearchResult{
                TmdbId: apiShows[1].Id,
                Title: apiShows[1].Name,
                CountriesEmoji: []string{"🇦🇺"},
                Year: 2022,
                Overview: apiShows[1].Overview,
                Slug: show2Slug,
                UrlPath: fmt.Sprintf("%d-%s", apiShows[1].Id, show2Slug),
            },
            ShowSearchResult{
                TmdbId: apiShows[2].Id,
                Title: apiShows[2].Name,
                CountriesEmoji: []string{"🇫🇷"},
                Year: 2018,
                Overview: apiShows[2].Overview,
                PosterImage: &show3PosterImage,
                Slug: show3Slug,
                UrlPath: fmt.Sprintf("%d-%s", apiShows[2].Id, show3Slug),
            },
            ShowSearchResult{
                TmdbId: apiShows[3].Id,
                Title: apiShows[3].Name,
                CountriesEmoji: []string{"🇮🇴"},
                Year: 0,
                Overview: apiShows[3].Overview,
                Slug: show4Slug,
                UrlPath: fmt.Sprintf("%d-%s", apiShows[3].Id, show4Slug),
            },
        },
        TotalResults: 4,
        TotalPages: 1,
    }

    result := suite.showSearch.Search(suite.ctx, searchToken, page)
    assert.Equal(suite.T(), expected, result)
}

func TestShowSearchTestSuite(t *testing.T) {
    suite.Run(t, new(ShowSearchTestSuite))
}
