package controllers

import (
    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
)

type HomeController struct {
    showSearch *show.ShowSearch
    trendingShows *show.TrendingShows
}

func NewHomeController(showSearch *show.ShowSearch, trendingShows *show.TrendingShows) *HomeController {
    return &HomeController{showSearch, trendingShows}
}

func (homeController *HomeController) Get(ctx *fiber.Ctx) error {
    if ctx.Query("q") != "" {
        return homeController.showSearchView(ctx)
    }

    return homeController.showHomeView(ctx)
}

func (homeController *HomeController) showHomeView(ctx *fiber.Ctx) error {
    shows := homeController.trendingShows.GetAll(ctx.Context())

    context := map[string]interface{}{
        "trending_shows": shows,
    }

    return RenderWithContext("home.html", context, ctx)
}

func (homeController *HomeController) showSearchView(ctx *fiber.Ctx) error {
    searchToken := ctx.Query("q")

    context := map[string]interface{}{
        "search_token": searchToken,
        "results": homeController.showSearch.Search(ctx.Context(), searchToken, 0),
    }

    return RenderWithContext("search.html", context, ctx)
}
