package db

import (
	"fmt"
)

var (
	// DefaultMaxOpenConnsNum : 默认的最大mysql连接数
	DefaultMaxOpenConnsNum = 20
	// DefaultMaxIdleConnsNum : 默认的最大mysql空闲连接数
	DefaultMaxIdleConnsNum = 0
)

// MysqlConfig : mysql 数据库配置
type MysqlConfig struct {
	MaxOpenConnsNum int // 最大打开连接数
	MaxIdleConnsNum int // 最大空闲连接数
	Addr            string
	Username        string
	Password        string
	DBName          string
}

func (c *MysqlConfig) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		c.Username, c.Password, c.Addr, c.DBName)
}
