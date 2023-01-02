package service

import (
	"context"

	"github.com/doduykhang/muse/pkg/config"
	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type CacheService interface {
	Set(key string, value string) error	
	Get(key string)	(string, error)
}

type redisService struct {
	redisClient *redis.Client
}

func (service *redisService) Set(key string, value string) error {
	err := service.redisClient.Set(ctx, key, value, 0).Err()
    	if err != nil {
		return err
    	}
	return nil
}

func (service *redisService) Get(key string) (string, error) {
	val, err := service.redisClient.Get(ctx, key).Result()
    	if err != nil {
		return "", err
    	}
	return val, nil
}

func NewCacheService () CacheService {
	return &redisService{
		redisClient: config.GetRedisClient(),
	}
}
