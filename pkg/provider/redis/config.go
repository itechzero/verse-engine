package redis

import (
	"time"

	"github.com/spf13/viper"
	"github.com/versegeek/toolkit/pkg/common"
)

type Config struct {
	enabled      bool
	Address      string
	DB           int
	UserName     string
	Password     string
	PoolSize     int
	MinIdleConns int
	MaxRetries   int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

const (
	defaultAddress      = "127.0.0.1:6379"
	defaultDB           = 0
	defaultUserName     = ""
	defaultPassword     = ""
	defaultPoolSize     = 50
	defaultMinIdleConns = 5
	defaultMaxRetries   = 3
)

var (
	defaultDialTimeout  = 3 * time.Second
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = defaultReadTimeout
)

func NewConfigFromEnv() *Config {
	v := viper.New()
	common.LoadFromFile(v)

	v.SetDefault("REDIS_ADDRESS", defaultAddress)
	v.SetDefault("REDIS_DB", defaultDB)
	v.SetDefault("REDIS_USERNAME", defaultUserName)
	v.SetDefault("REDIS_PASSWORD", defaultPassword)
	v.SetDefault("REDIS_POOL_SIZE", defaultPoolSize)
	v.SetDefault("REDIS_MIN_IDLE_CONNS", defaultMinIdleConns)
	v.SetDefault("REDIS_MAX_RETRIES", defaultMaxRetries)
	v.SetDefault("REDIS_DIAL_TIMEOUT", defaultDialTimeout)
	v.SetDefault("REDIS_READ_TIMEOUT", defaultReadTimeout)
	v.SetDefault("REDIS_WRITE_TIMEOUT", defaultWriteTimeout)

	return &Config{
		Address:      v.GetString("REDIS_ADDRESS"),
		DB:           v.GetInt("REDIS_DB"),
		UserName:     v.GetString("REDIS_USERNAME"),
		Password:     v.GetString("REDIS_PASSWORD"),
		PoolSize:     v.GetInt("REDIS_POOL_SIZE"),
		MinIdleConns: v.GetInt("REDIS_MIN_IDLE_CONNS"),
		MaxRetries:   v.GetInt("REDIS_MAX_RETRIES"),
		DialTimeout:  v.GetDuration("REDIS_DIAL_TIMEOUT"),
		ReadTimeout:  v.GetDuration("REDIS_READ_TIMEOUT"),
		WriteTimeout: v.GetDuration("REDIS_WRITE_TIMEOUT"),
	}
}
