package middleware

import (
    "regexp"
    "strconv"

    "github.com/gofiber/fiber/v2"

    "github.com/mikeyclarke/can-i-binge-yet/src/go/show"
)

type ShowLoaderMiddleware struct {
    showLoader show.ShowLoaderInterface
}

func NewShowLoaderMiddleware(showLoader show.ShowLoaderInterface) *ShowLoaderMiddleware {
    return &ShowLoaderMiddleware{showLoader}
}

func (middleware *ShowLoaderMiddleware) LoadShow(ctx *fiber.Ctx) error {
    showUrlPath := ctx.Params("show_url_path")
    re := regexp.MustCompile("^([0-9]+)-")

    matches := re.FindStringSubmatch(showUrlPath)
    if matches == nil || len(matches) < 2 || matches[1] == "" {
        return fiber.ErrNotFound
    }

    id, err := strconv.Atoi(matches[1])
    if err != nil {
        return err
    }

    slugIndex := len(matches[1]) + 1
    slug := showUrlPath[slugIndex:]

    if slug == "" {
        return fiber.ErrNotFound
    }

    show := middleware.showLoader.Load(ctx.Context(), id)
    if show == nil {
        return fiber.ErrNotFound
    }

    if slug != show.Slug {
        return ctx.RedirectToRoute("front_end.app.show", fiber.Map{
            "show_url_path": show.UrlPath,
        })
    }

    ctx.Locals("show", show)

    return ctx.Next()
}
