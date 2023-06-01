package redisclient

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"lion/internal/shared/config"
	"time"
)

type Redis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Exist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
}

type DefaultClient struct {
	config *config.Config
	client *redis.Client
}

func NewDefaultClient(config *config.Config) Redis {
	addr := config.Storage.Redis.Address
	db := config.Storage.Redis.DB

	opt := redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              "",
		CredentialsProvider:   nil,
		DB:                    db,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	}
	redisClient := redis.NewClient(&opt)

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println("failed to connect redis:", err)
	}

	return &DefaultClient{
		config: config,
		client: redisClient,
	}
}

func (d DefaultClient) Get(ctx context.Context, key string) (string, error) {
	return d.client.Get(ctx, key).Result()
}

func (d DefaultClient) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return d.client.Set(ctx, key, value, ttl).Err()
}

func (d DefaultClient) Exist(ctx context.Context, key string) bool {
	exists, _ := d.client.Exists(ctx, key).Result()
	return exists == 1
}

func (d DefaultClient) Delete(ctx context.Context, key string) error {
	return d.client.Del(ctx, key).Err()
}
