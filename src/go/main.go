package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    templateRenderer := CreateTemplateRenderer()
    showLoaderMiddleware := CreateShowLoaderMiddleware()

    app := fiber.New(fiber.Config{
        Views: templateRenderer,
    })

    homeController := CreateHomeController()
    showController := CreateShowController()
    seasonEpisodesController := CreateSeasonEpisodesController()

    showPattern := "/show/:show_url_path"

    app.Static("/", "public")
    app.Get("/", homeController.Get).Name("front_end.app.home")
    app.Use(showPattern, showLoaderMiddleware.LoadShow)
    app.Get(showPattern, showController.Get).Name("front_end.app.show")
    app.Get("/_/show/:show_id/season-episodes/:season_number", seasonEpisodesController.Get).Name("api.app.season_episodes.get")

    log.Fatal(app.Listen(":8000"))
}
