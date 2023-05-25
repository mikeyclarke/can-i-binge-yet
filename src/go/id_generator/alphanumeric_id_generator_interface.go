package id_generator

type AlphanumericIdGeneratorInterface interface {
    Generate(length int) string
}
