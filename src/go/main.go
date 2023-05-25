package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    templateRenderer := CreateTemplateRenderer()

    app := fiber.New(fiber.Config{
        Views: templateRenderer,
    })

    homeController := CreateHomeController()

    app.Static("/", "public")
    app.Get("/", homeController.Get).Name("front_end.app.home")

    log.Fatal(app.Listen(":8000"))
}
