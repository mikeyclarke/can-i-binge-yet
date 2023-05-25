package cache

import (
    "context"
    "time"

    redisCache "github.com/go-redis/cache/v8"
)

type Cache struct {
    redisCache *redisCache.Cache
}

func NewCache(redisCache *redisCache.Cache) *Cache {
    return &Cache{redisCache}
}

func (cache *Cache) Get(ctx context.Context, key string, result interface{}) (bool, error) {
    var exists bool
    err := cache.redisCache.Get(ctx, key, result)

    if err == nil {
        exists = true
    }

    if err == redisCache.ErrCacheMiss {
        return exists, nil
    }

    return exists, err
}

func (cache *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    return cache.redisCache.Set(&redisCache.Item{
        Ctx: ctx,
        Key: key,
        Value: value,
        TTL: ttl,
    })
}
