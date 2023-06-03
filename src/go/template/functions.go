package template

import (
    "errors"
    "fmt"
    "net/url"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/nikolalohinski/gonja/exec"
)

var Functions = exec.NewContext(map[string]interface{}{
    "icon": functionIcon,
    "return_to_url": functionReturnToUrl,
})

func functionReturnToUrl(va *exec.VarArgs) *exec.Value {
    args := va.ExpectArgs(2)

    ctx, okay := args.Args[0].Interface().(*fiber.Ctx)
    if !okay {
        return exec.AsValue(errors.New("first argument to “return_to_url” function should be a fiber.Ctx instance"))
    }

    if !args.Args[1].IsString() {
        return exec.AsValue(errors.New("second argument to “return_to_url” function should be a string"))
    }
    fallbackUrl := exec.AsValue(args.Args[1].String())

    returnTo := ctx.Query("return_to")
    if returnTo == "" {
        return fallbackUrl
    }

    parsedUrl, err := url.Parse(returnTo)
    if err != nil {
        return fallbackUrl
    }

    if parsedUrl.Scheme != "" || parsedUrl.Host != "" || !strings.HasPrefix(parsedUrl.Path, "/") {
        return fallbackUrl
    }

    return exec.AsValue(returnTo)
}

func functionIcon(va *exec.VarArgs) *exec.Value {
    args := va.Expect(2, []*exec.KwArg{
        &exec.KwArg{Name: "label", Default: nil},
    })

    if !args.Args[0].IsString() {
        return exec.AsValue(errors.New("first argument to “icon” function should be a string"))
    }

    if !args.Args[1].IsString() {
        return exec.AsValue(errors.New("second argument to “icon” function should be a string"))
    }

    name := args.Args[0].String()
    className := args.Args[1].String()
    labelArg := args.KwArgs["label"]

    var ariaAttribute string
    if labelArg.IsNil() {
        ariaAttribute = "aria-hidden=\"true\""
    } else {
        label := labelArg.String()
        ariaAttribute = fmt.Sprintf("aria-label=\"%s\"", label)
    }

    result := fmt.Sprintf(
        "<svg class=\"%s\" %s><use xlink:href=\"#icon-sprite__%s\"/></svg>",
        className,
        ariaAttribute,
        name,
    )
    return exec.AsSafeValue(result)
}
