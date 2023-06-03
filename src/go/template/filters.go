package template

import (
    "errors"
    "fmt"
    "time"

    "github.com/goodsign/monday"
    "github.com/nikolalohinski/gonja/exec"
)

var Filters = map[string]exec.FilterFunction{
    "format_date": filterFormatDate,
}

func filterFormatDate(e *exec.Evaluator, in *exec.Value, params *exec.VarArgs) *exec.Value {
    if in.IsError() {
        return in
    }

    p := params.Expect(0, []*exec.KwArg{
        &exec.KwArg{Name: "format", Default: "medium"},
        &exec.KwArg{Name: "locale", Default: "en_US"},
    })

    if in.IsNil() {
        return in
    }
    t, okay := in.Interface().(*time.Time)
    if !okay {
        return exec.AsValue(errors.New("first argument to “format_date” filter should be a time.Time instance"))
    }

    l := p.KwArgs["locale"].String()
    var locale monday.Locale
    for _, candidate := range monday.ListLocales() {
        if string(candidate) == l {
            locale = candidate
            break
        }
    }
    if locale == "" {
        return exec.AsValue(errors.New("unsupported locale provided for “format_date” filter"))
    }

    f := p.KwArgs["format"]
    layoutStyle := f.String()

    var layout string
    switch layoutStyle {
        case "short":
            layout = monday.ShortFormatsByLocale[locale]
        case "medium":
            layout = monday.MediumFormatsByLocale[locale]
        case "long":
            layout = monday.LongFormatsByLocale[locale]
        case "full":
            layout = monday.FullFormatsByLocale[locale]
        default:
            return exec.AsValue(fmt.Errorf(
                "unsupported value for keyword argument “format” in “format_date” filter; " +
                "supported values are %s, %s provided",
                "“short”, “medium”, “long”, and “full”",
                layoutStyle,
            ))
    }

    result := monday.Format(*t, layout, locale)
    return exec.AsSafeValue(result)
}
