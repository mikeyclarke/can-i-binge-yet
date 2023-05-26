package show

import (
    "context"
    "fmt"
    "strings"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/themoviedb"
)

type ShowImageFormatter struct {
    tmdbConfig themoviedb.TheMovieDbConfigurationInterface
}

func NewShowImageFormatter(tmdbConfig themoviedb.TheMovieDbConfigurationInterface) *ShowImageFormatter {
    return &ShowImageFormatter{tmdbConfig}
}

func (formatter *ShowImageFormatter) Format(
    ctx context.Context,
    imageType string,
    imagePath string,
    defaultSize string,
) ShowImageResult {
    if defaultSize == "" {
        defaultSize = "largest"
    }
    imagesBaseUrl := formatter.tmdbConfig.GetImageBaseUrl(ctx)

    var defaultSrc string
    var srcsets []string

    sizes := formatter.tmdbConfig.GetImageSizes(ctx, imageType)
    var relevantSizes []string
    for _, size := range sizes {
        if strings.HasPrefix(size, "w") {
            relevantSizes = append(relevantSizes, size)
        }
    }

    loopMax := len(relevantSizes) - 1
    for index, size := range relevantSizes {
        src := fmt.Sprintf("%s%s%s", imagesBaseUrl, size, imagePath)
        sizePixels := size[1:]
        srcsets = append(srcsets, fmt.Sprintf("%s %sw", src, sizePixels))

        if (index == 0 && defaultSize == "smallest") || (index == loopMax && defaultSize == "largest") {
            defaultSrc = src
        }
    }

    return ShowImageResult{defaultSrc, strings.Join(srcsets, ", ")}
}
