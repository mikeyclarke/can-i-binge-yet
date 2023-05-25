package show

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/mocks/context"
)

func Test_Format(t *testing.T) {
    tmdbConfig := themoviedb.NewMockTheMovieDbConfigurationInterface(t)
    formatter := NewShowImageFormatter(tmdbConfig)

    ctx := context.NewMockContext(t)
    imageType := "poster"
    imagePath := "/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg"
    defaultSize := ""

    imageBaseUrl := "https://image.tmdb.org/t/p/"
    sizes := []string{
        "w92",
        "w154",
        "w185",
        "w342",
        "w500",
        "w780",
        "original",
    }

    tmdbConfig.On("GetImageBaseUrl", ctx).Return(imageBaseUrl)
    tmdbConfig.On("GetImageSizes", ctx, imageType).Return(sizes)

    srcsets := []string{
        "https://image.tmdb.org/t/p/w92/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 92w",
        "https://image.tmdb.org/t/p/w154/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 154w",
        "https://image.tmdb.org/t/p/w185/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 185w",
        "https://image.tmdb.org/t/p/w342/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 342w",
        "https://image.tmdb.org/t/p/w500/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 500w",
        "https://image.tmdb.org/t/p/w780/bfxwMdQyJc0CL24m5VjtWAN30mt.jpg 780w",
    }

    expected := ShowImageResult{
        strings.Split(srcsets[len(srcsets)-1], " ")[0],
        strings.Join(srcsets, ", "),
    }
    result := formatter.Format(ctx, imageType, imagePath, defaultSize)
    assert.Equal(t, expected, result)
}
