package main

import (
    "errors"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/front_end/app/controllers"
)

func main() {
    templateRenderer := CreateTemplateRenderer()
    showLoaderMiddleware := CreateShowLoaderMiddleware()

    app := fiber.New(fiber.Config{
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            var e *fiber.Error
            if errors.As(err, &e) {
                code = e.Code
            }

            if code == 404 || code == 500 {
                err = controllers.RenderWithContext("errorpage.html", map[string]interface{}{
                    "code": code,
                }, ctx)
            }

            if err != nil {
                return ctx.Status(code).SendString("Internal Server Error")
            }

            return nil
        },
        Views: templateRenderer,
    })
    recoverConfig := recover.Config{
        Next: nil,
        EnableStackTrace: true,
    }
    app.Use(recover.New(recoverConfig))

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
