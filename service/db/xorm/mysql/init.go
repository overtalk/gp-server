package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// MysqlDriver : mysql driver
type MysqlDriver struct {
	config *MysqlConfig
	conn   *xorm.Engine
}

// NewMysqlDriver : constructor of MysqlDriver
func NewMysqlDriver(config *MysqlConfig) (*MysqlDriver, error) {
	// Set the default settings
	if config.MaxOpenConnsNum <= 0 {
		config.MaxIdleConnsNum = DefaultMaxOpenConnsNum
	}
	if config.MaxIdleConnsNum <= 0 {
		config.MaxIdleConnsNum = DefaultMaxIdleConnsNum
	}

	mysqlDriver := &MysqlDriver{
		config: config,
	}

	if err := mysqlDriver.Connect(); err != nil {
		return nil, err
	}

	return mysqlDriver, nil
}

// Connect : connect to the database and initialize the database connection
func (m *MysqlDriver) Connect() error {
	if m.config == nil {
		return ErrNoMysqlConf
	}

	db, err := xorm.NewEngine("mysql", m.config.getDSN())
	if err != nil {
		return err
	}
	m.conn = db

	// initialize the database connection
	m.initialization()

	return nil
}

func (m *MysqlDriver) initialization() {
	// set the number of idle pools in the connection pool
	m.conn.SetMaxIdleConns(m.config.MaxIdleConnsNum)
	// set the number of maximum open connections
	m.conn.SetMaxOpenConns(m.config.MaxOpenConnsNum)

	if m.config.IsCached {
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		m.conn.SetDefaultCacher(cacher)
	}
}
