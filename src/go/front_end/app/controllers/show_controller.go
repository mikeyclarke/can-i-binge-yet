package controllers

import (
    "errors"
    "strings"

    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
)

type ShowController struct {}

func NewShowController() *ShowController {
    return &ShowController{}
}

func (showController *ShowController) Get(ctx *fiber.Ctx) error {
    s := ctx.Locals("show")
    typedShow, ok := s.(*show.ShowPageResult)
    if !ok {
        return errors.New("unexpected value type")
    }

    context := map[string]interface{}{
        "html_title": typedShow.Title,
        "show": typedShow,
        "showMeta": showController.formatShowMeta(typedShow),
    }

    return RenderWithContext("show.html", context, ctx)
}

func (showController *ShowController) formatShowMeta(show *show.ShowPageResult) string {
    result := []string{}

    if show.AirDatesDisplay != "" {
        result = append(result, show.AirDatesDisplay)
    }

    if len(show.CountriesEmoji) > 0 {
        result = append(result, strings.Join(show.CountriesEmoji, " "))
    }

    return strings.Join(result, " ")
}
