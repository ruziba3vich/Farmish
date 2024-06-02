package redis

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisDB(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: "",
		DB:       0,
	})
	return rdb, nil
}
