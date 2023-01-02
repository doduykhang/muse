package config

import (
    "github.com/go-redis/redis/v9"
)

var (
	redisClient *redis.Client
)

func init() {
    redisClient = redis.NewClient(&redis.Options{
	Addr:     env.RedisHost + ":" + env.RedisPort,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}

func GetRedisClient () *redis.Client {
	return redisClient
}
