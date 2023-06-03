package controllers

import (
    "strconv"

    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
    "github.com/mikeyclarke/can-i-binge-yet/src/go/template"
)

type SeasonEpisodesController struct {
    seasonEpisodesLoader show.SeasonEpisodesLoaderInterface
    templateRenderer *template.TemplateRenderer
}

func NewSeasonEpisodesController(
    seasonEpisodesLoader show.SeasonEpisodesLoaderInterface,
    templateRenderer *template.TemplateRenderer,
) *SeasonEpisodesController {
    return &SeasonEpisodesController{seasonEpisodesLoader, templateRenderer}
}

func (controller *SeasonEpisodesController) Get(ctx *fiber.Ctx) error {
    showIdParam := ctx.Params("show_id")
    seasonNumberParam := ctx.Params("season_number")

    showId, err := strconv.Atoi(showIdParam)
    if err != nil {
        return err
    }

    seasonNumber, err := strconv.Atoi(seasonNumberParam)
    if err != nil {
        return err
    }

    episodes := controller.seasonEpisodesLoader.Load(showId, seasonNumber)
    if episodes == nil {
        return fiber.ErrNotFound
    }

    html, err := controller.templateRenderer.RenderToString("show/_season_episodes.html", map[string]interface{}{
        "ctx": ctx,
        "season": map[string]interface{}{
            "Episodes": episodes,
        },
    })
    if err != nil {
        return err
    }

    return ctx.JSON(fiber.Map{
        "html": html,
    })
}
