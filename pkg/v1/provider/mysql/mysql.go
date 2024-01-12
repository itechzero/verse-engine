package mysql

import (
	"github.com/itechzero/lib-core-go/pkg/v1/provider"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ provider.Provider = new(MySQL)

type MySQL struct {
	config *Config
	conn   *gorm.DB
}

func New(config *Config) *MySQL {
	if config == nil {
		config = NewConfig()
	}

	return &MySQL{
		config: config,
	}
}

func (p *MySQL) Enabled() bool {
	return p.config.enabled
}

func (p *MySQL) Init() error {
	db, err := gorm.Open(mysql.Open(p.config.dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("[MySQL] Connect failed")
		return err
	}

	p.conn = db

	logrus.WithField("dsn", p.DSN()).Debug("[MySQL] initialized")

	return nil
}

func (p *MySQL) Close() error {
	db, _ := p.conn.DB()
	return db.Close()
}

func (p *MySQL) DSN() string {
	return p.config.dsn
}
