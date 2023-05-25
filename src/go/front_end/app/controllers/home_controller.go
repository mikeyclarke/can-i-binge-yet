package controllers

import (
    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
)

type HomeController struct {
    trendingShows *show.TrendingShows
}

func NewHomeController(trendingShows *show.TrendingShows) *HomeController {
    return &HomeController{trendingShows}
}

func (homeController *HomeController) Get(ctx *fiber.Ctx) error {
    shows := homeController.trendingShows.GetAll(ctx.Context())

    context := map[string]interface{}{
        "trending_shows": shows,
    }

    return RenderWithContext("home.html", context, ctx)
}
