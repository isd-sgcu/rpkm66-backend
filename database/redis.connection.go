package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
)

func InitRedisConnect(conf *cfgldr.Redis) (cache *redis.Client, err error) {
	cache = redis.NewClient(&redis.Options{
		Addr: conf.Host,
		DB:   3,
	})

	return
}
