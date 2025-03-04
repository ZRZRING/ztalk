package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"ztalk/pkg/settings"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	return rdb.Ping().Err()
}

func Close() {
	_ = rdb.Close()
}
