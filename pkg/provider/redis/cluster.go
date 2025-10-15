package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/versegeek/toolkit/pkg/provider"
)

var _ provider.Provider = new(Cluster)

type Cluster struct {
	client *redis.ClusterClient
	config *Config
}

func NewCluster(config *Config) *Cluster {
	if config == nil {
		config = NewConfigFromEnv()
	}

	return &Cluster{
		config: config,
	}
}

func (c *Cluster) ClusterClient() *redis.ClusterClient {
	return c.client
}

func (c *Cluster) Enabled() bool {
	return c.config.enabled
}

func (c *Cluster) Init() error {
	c.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{c.config.Address},
		Username:     c.config.UserName,
		Password:     c.config.Password,
		MaxRetries:   c.config.MaxRetries,
		DialTimeout:  c.config.DialTimeout,
		ReadTimeout:  c.config.ReadTimeout,
		WriteTimeout: c.config.WriteTimeout,
		PoolSize:     c.config.PoolSize,
		MinIdleConns: c.config.MinIdleConns,
	})

	return c.client.Ping(context.Background()).Err()
}

func (c *Cluster) Close() error {
	return c.client.Close()
}
