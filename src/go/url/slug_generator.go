package url

import (
    "log"

    "github.com/gosimple/slug"
)

type SlugGenerator struct {}

func NewSlugGenerator() *SlugGenerator {
    return &SlugGenerator{}
}

func (slugGenerator *SlugGenerator) Generate(textToSlugify string, maxLength int) string {
    if maxLength == 0 {
        maxLength = 120
    }

    slug.MaxLength = maxLength
    result := slug.Make(textToSlugify)
    if len(result) == 0 {
        log.Fatalf("%s cannot be slugified", textToSlugify)
    }

    return result
}
