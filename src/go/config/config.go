package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/ilyakaznacheev/cleanenv"
    "github.com/joho/godotenv"
)

type AssetDirectory string
type AssetManifestPath string
type RedisHost string
type RedisPort int
type RedisPassword string
type SearchInputPlaceholderExample string
type Snowfall bool
type TemplatesDirectory string
type TmdbApiBaseUrl string
type TmdbApiKey string

func EnsureNonEmptyString(s string) error {
    if s == "" {
        return fmt.Errorf("field value can't be empty")
    }
    return nil
}

func StrToBool(s string) (bool, error) {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return false, err
    }

    return strconv.ParseBool(s)
}

func StrToInt(s string) (int, error) {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return 0, err
    }

    return strconv.Atoi(s)
}

func (f *AssetDirectory) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = AssetDirectory(s)
    return nil
}

func (f *AssetManifestPath) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = AssetManifestPath(s)
    return nil
}

func (f *RedisHost) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = RedisHost(s)
    return nil
}

func (f *RedisPort) SetValue(s string) error {
    i, err := StrToInt(s)
    if err != nil {
        return err
    }
    *f = RedisPort(i)
    return nil
}

func (f *RedisPassword) SetValue(s string) error {
    *f = RedisPassword(s)
    return nil
}

func (f *SearchInputPlaceholderExample) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = SearchInputPlaceholderExample(s)
    return nil
}

func (f *Snowfall) SetValue(s string) error {
    b, err := StrToBool(s)
    if err != nil {
        return err
    }
    *f = Snowfall(b)
    return nil
}

func (f *TemplatesDirectory) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = TemplatesDirectory(s)
    return nil
}

func (f *TmdbApiBaseUrl) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = TmdbApiBaseUrl(s)
    return nil
}

func (f *TmdbApiKey) SetValue(s string) error {
    err := EnsureNonEmptyString(s)
    if err != nil {
        return err
    }
    *f = TmdbApiKey(s)
    return nil
}

type ApplicationConfig struct {
    AssetDirectory AssetDirectory `yaml:"asset_directory"`
    AssetManifestPath AssetManifestPath `yaml:"asset_manifest_path"`
    RedisHost RedisHost `env:"REDIS_HOST" env-required:"true"`
    RedisPort RedisPort `env:"REDIS_PORT" env-required:"true"`
    RedisPassword RedisPassword `env:"REDIS_PASSWORD" env-required:"true"`
    SearchInputPlaceholderExample SearchInputPlaceholderExample `yaml:"search_input_placeholder_example"`
    Snowfall Snowfall `yaml:"snowfall"`
    TemplatesDirectory TemplatesDirectory `yaml:"templates_directory"`
    TmdbApiBaseUrl TmdbApiBaseUrl `yaml:"tmdb_api_base_url"`
    TmdbApiKey TmdbApiKey `env:"TMDB_API_KEY" env-required:"true"`
}

var cfg ApplicationConfig

func NewApplicationConfig() ApplicationConfig {
    envType := os.Getenv("APP_ENV")
    if envType == "dev" {
        err := godotenv.Load()
        if err != nil {
            panic(err)
        }
    }

    err := cleanenv.ReadConfig("config/config.yaml", &cfg)
    if err != nil {
        panic(err)
    }
    return cfg
}
