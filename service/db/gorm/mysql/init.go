package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/qinhan-shu/gp-server/model/gorm"
)

// MysqlDriver : mysql 驱动
type MysqlDriver struct {
	config *MysqlConfig
	conn   *gorm.DB
}

// NewMysqlDriver : MysqlDriver 的构造方法
func NewMysqlDriver(config *MysqlConfig) (*MysqlDriver, error) {
	// 设置默认连接数量
	if config.MaxOpenConnsNum <= 0 {
		config.MaxIdleConnsNum = DefaultMaxOpenConnsNum
	}

	// 设置默认空闲连接数量
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

// Connect : 连接数据库，并初始化数据库连接
func (m *MysqlDriver) Connect() error {
	if m.config == nil {
		return ErrNoMysqlConf
	}

	db, err := gorm.Open("mysql", m.config.getDSN())
	if err != nil {
		return err
	}
	m.conn = db

	// 初始化数据库连接（数据库最多连接数量，自动迁移，设置表名称...）
	m.initialization()

	return nil
}

func (m *MysqlDriver) initialization() {
	//连接池的空闲数大小
	m.conn.DB().SetMaxIdleConns(m.config.MaxIdleConnsNum)
	//最大打开连接数
	m.conn.DB().SetMaxIdleConns(m.config.MaxOpenConnsNum)

	// 设置表名就是结构体的名字
	// 如果不设置的话，表名默认为结构体名的复数
	m.conn.SingularTable(true)

	// 自动迁移，对于注册的model， 将增加数据库中没有但是model中定义过的字段，不会删除（改变）原先的字段&数据
	//m.conn.AutoMigrate(&model.Test{})
}
