package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/itechzero/verse-engine/pkg/v1/provider"
)

var _ provider.Provider = new(Cluster)

type Cluster struct {
	clientCluster *redis.ClusterClient
	config        *Config
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
	return c.clientCluster
}

func (c *Cluster) Enabled() bool {
	return c.config.enabled
}

func (c *Cluster) Init() error {
	c.clientCluster = redis.NewClusterClient(&redis.ClusterOptions{
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

	return c.clientCluster.Ping(context.Background()).Err()
}

func (c *Cluster) Close() error {
	return c.clientCluster.Close()
}
