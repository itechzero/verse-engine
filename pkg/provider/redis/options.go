package redis

import "github.com/redis/go-redis/v9"

type (
	Options struct{}

	Option func(*redis.Options)
)
