package mysql

import (
	"time"

	"github.com/itechzero/lib-core-go/pkg/v1/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DSNKey = "MYSQL_DSN"

	HostKey     = "MYSQL_HOST"
	UserKey     = "MYSQL_USER"
	PwdKey      = "MYSQL_PASSWORD"
	DatabaseKey = "MYSQL_DATABASE"

	MaxIdleKey          = "MYSQL_MAX_IDLE_CONNS"
	MaxOpenKey          = "MYSQL_MAX_OPEN_CONNS"
	MaxLifetimeKey      = "MYSQL_MAX_LIFETIME_MINUTES"
	TimeoutKey          = "MYSQL_TIMEOUT_SECOND"
	MigrationKey        = "MYSQL_MIGRATION"
	MigrateDirectoryKey = "MYSQL_MIGRATE_DIRECTORY"
)

type Config struct {
	enabled  bool
	dsn      string
	Host     string
	User     string
	Pwd      string
	Database string

	MaxOpenConns     int
	MaxIdleConns     int
	MaxLifetime      time.Duration
	Timeout          time.Duration
	Migration        string
	MigrateDirectory string
}

func NewConfig() *Config {
	v := viper.New()

	v.SetDefault(HostKey, "127.0.0.1:3306")
	v.SetDefault(DatabaseKey, "")
	v.SetDefault(UserKey, "root")
	v.SetDefault(PwdKey, "")

	v.SetDefault(MaxIdleKey, 30)
	v.SetDefault(MaxOpenKey, 30)
	v.SetDefault(MaxLifetimeKey, 60)
	v.SetDefault(TimeoutKey, 10)
	v.SetDefault(MigrationKey, "")
	v.SetDefault(MigrateDirectoryKey, "./scripts/sql")

	common.LoadFromFile(v)

	config := &Config{
		dsn: v.GetString(DSNKey),

		Host:     v.GetString(HostKey),
		User:     v.GetString(UserKey),
		Pwd:      v.GetString(PwdKey),
		Database: v.GetString(DatabaseKey),

		MaxOpenConns:     v.GetInt(MaxOpenKey),
		MaxIdleConns:     v.GetInt(MaxIdleKey),
		MaxLifetime:      v.GetDuration(MaxLifetimeKey) * time.Minute,
		Timeout:          v.GetDuration(TimeoutKey) * time.Second,
		Migration:        v.GetString(MigrationKey),
		MigrateDirectory: v.GetString(MigrateDirectoryKey),
	}

	logrus.WithFields(logrus.Fields{
		"max_open_conns": config.MaxOpenConns,
		"max_idle_conns": config.MaxIdleConns,
		"max_lifetime":   config.MaxLifetime,
		"timeout_second": config.Timeout.Seconds(),
	}).Debug("[Mysql] Config initialized")

	return config
}
