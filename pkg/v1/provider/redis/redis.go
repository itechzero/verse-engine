package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/itechzero/lib-core-go/pkg/v1/provider"
)

var _ provider.Provider = new(Redis)

type Redis struct {
	client *redis.Client
	config *Config
}

func New(config *Config) *Redis {
	if config == nil {
		config = NewConfigFromEnv()
	}

	return &Redis{
		config: config,
	}
}

func (r *Redis) Client() *redis.Client {
	return r.client
}

func (r *Redis) Enabled() bool {
	return r.config.enabled
}

func (r *Redis) Init() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:         r.config.Address,
		Username:     r.config.UserName,
		Password:     r.config.Password,
		DB:           r.config.DB,
		MaxRetries:   r.config.MaxRetries,
		DialTimeout:  r.config.DialTimeout,
		ReadTimeout:  r.config.ReadTimeout,
		WriteTimeout: r.config.WriteTimeout,
		PoolSize:     r.config.PoolSize,
		MinIdleConns: r.config.MinIdleConns,
	})

	return r.client.Ping(context.Background()).Err()
}

func (r *Redis) Close() error {
	return r.client.Close()
}
