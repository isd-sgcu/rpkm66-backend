package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
)

func InitRedisConnect(conf *config.Redis) (cache *redis.Client, err error) {
	cache = redis.NewClient(&redis.Options{
		Addr: conf.Host,
		DB:   3,
	})

	return
}
