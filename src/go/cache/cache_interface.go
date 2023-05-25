package cache

import (
    "context"
    "time"
)

type CacheInterface interface {
    Get(ctx context.Context, key string, result interface{}) (bool, error)
    Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}
