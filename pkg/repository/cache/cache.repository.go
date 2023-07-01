package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/isd-sgcu/rpkm66-backend/internal/repository/cache"
)

type Repository interface {
	RemoveCache(key string) error
	SaveCache(key string, value interface{}, ttl int) error
	GetCache(key string, value interface{}) error
}

func NewRepository(client *redis.Client) Repository {
	return cache.NewRepository(client)
}
