package common_client

import (
	"time"

	"github.com/go-redis/redis/v8"
	"{{.Name}}/src/common/boot_config"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:               boot_config.BootConfigData.Redis.Addr,
		Password:           boot_config.BootConfigData.Redis.Password,
		DB:                 boot_config.BootConfigData.Redis.DB,
		MaxRetries:         2,
		DialTimeout:        time.Second * 10,
		ReadTimeout:        time.Second * 5,
		WriteTimeout:       time.Second * 5,
		PoolTimeout:        time.Second * 10,
		IdleTimeout:        time.Minute * 10,
		IdleCheckFrequency: time.Second * 30,
	})
	return &RedisClient{
		Client: client,
	}
}
