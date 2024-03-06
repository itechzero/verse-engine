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

func (m *MySQL) Enabled() bool {
	return m.config.enabled
}

func (m *MySQL) Init() error {
	db, err := gorm.Open(mysql.Open(m.config.dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("[MySQL] Connect failed")
		return err
	}

	m.conn = db
	logrus.WithField("dsn", m.config.dsn).Debug("[MySQL] initialized")

	return nil
}

func (m *MySQL) Close() error {
	db, _ := m.conn.DB()
	return db.Close()
}

func (m *MySQL) Client() *gorm.DB {
	return m.conn
}
