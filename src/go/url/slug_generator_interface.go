package url

type SlugGeneratorInterface interface {
    Generate(textToSlugify string, maxLength int) string
}
