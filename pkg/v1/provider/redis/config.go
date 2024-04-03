package redis

import (
	"time"

	"github.com/itechzero/verse-engine/pkg/v1/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	address := v.GetString("REDIS_ADDRESS")
	db := v.GetInt("REDIS_DB")
	userName := v.GetString("REDIS_USERNAME")
	password := v.GetString("REDIS_PASSWORD")
	poolSize := v.GetInt("REDIS_POOL_SIZE")
	minIdleConns := v.GetInt("REDIS_MIN_IDLE_CONNS")
	maxRetries := v.GetInt("REDIS_MAX_RETRIES")
	dialTimeout := v.GetDuration("REDIS_DIAL_TIMEOUT")
	readTimeout := v.GetDuration("REDIS_READ_TIMEOUT")
	writeTimeout := v.GetDuration("REDIS_WRITE_TIMEOUT")

	logrus.WithFields(logrus.Fields{
		"address":        address,
		"db":             db,
		"user_name":      userName,
		"password":       password,
		"pool_size":      poolSize,
		"min_idle_conns": minIdleConns,
		"max_retries":    maxRetries,
		"dial_timeout":   dialTimeout,
		"read_timeout":   readTimeout,
		"write_timeout":  writeTimeout,
	}).Debugf("REDIS Config initialized")

	return &Config{
		Address:      address,
		DB:           db,
		UserName:     userName,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		MaxRetries:   maxRetries,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}
