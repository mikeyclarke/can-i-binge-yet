package controllers

import (
    "github.com/gofiber/fiber/v2"
)

func RenderWithContext(templateFilename string, templateContext map[string]interface{}, ctx *fiber.Ctx) error {
    templateContext["ctx"] = ctx
    return ctx.Render(templateFilename, templateContext)
}
