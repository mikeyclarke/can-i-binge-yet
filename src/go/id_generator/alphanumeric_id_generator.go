package id_generator

import (
    "math/rand"
)

type AlphanumericIdGenerator struct {}

func NewAlphanumericIdGenerator() *AlphanumericIdGenerator {
    return &AlphanumericIdGenerator{}
}

func (generator *AlphanumericIdGenerator) Generate(length int) string {
    const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    result := make([]byte, length)
    for i := range result {
        result[i] = characters[rand.Int63() % int64(len(characters))]
    }
    return string(result)
}
