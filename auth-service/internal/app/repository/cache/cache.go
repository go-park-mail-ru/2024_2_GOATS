package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICached interface {
	GetToken(ctx context.Context, token string) (string, error)
	SetToken(ctx context.Context, token string, data string, expiration time.Duration) error
	Close() error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) SetToken(ctx context.Context, token string, data string, expiration time.Duration) error {
	err := r.client.Set(ctx, token, data, expiration).Err()
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to set token in cache")
	}
	err = r.client.Set(ctx, data, token, expiration).Err()
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to set token in cache")
	}
	return nil
}

func (r *RedisCache) GetToken(ctx context.Context, email string) (string, error) {
	data, err := r.client.Get(ctx, email).Result()
	fmt.Println("datadata", data)
	if err == redis.Nil {
		return "", errors.New("token not found")
	} else if err != nil {
		return "", errors.New("failed to get token from cache")
	}
	return data, nil
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
