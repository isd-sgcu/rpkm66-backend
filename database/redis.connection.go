package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
)

func InitRedisConnect(conf *cfgldr.Redis) (cache *redis.Client, err error) {
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.Dbnum,
	})

	return
}
