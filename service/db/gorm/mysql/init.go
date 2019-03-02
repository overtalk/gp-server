package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/qinhan-shu/gp-server/model/gorm"
)

// MysqlDriver : mysql driver
type MysqlDriver struct {
	config *MysqlConfig
	conn   *gorm.DB
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

	db, err := gorm.Open("mysql", m.config.getDSN())
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
	m.conn.DB().SetMaxIdleConns(m.config.MaxIdleConnsNum)
	// set the number of maximum open connections
	m.conn.DB().SetMaxIdleConns(m.config.MaxOpenConnsNum)

	// set the table name is the name of the structure
	// is not, the table name be the plural number of structure name
	m.conn.SingularTable(true)

	// Automatic migration
	// for registered models, will increase the fields in the database but not defined in the model
	// will not delete (change) the original fields & data
	//m.conn.AutoMigrate(&model.Test{})
}
