package db

import (
	"fmt"
)

// MysqlConfig : mysql 数据库配置
type MysqlConfig struct {
	MaxConnection int // 最大连接数(暂时没什么用)
	Addr          string
	Username      string
	Password      string
	DBName        string
}

func (c *MysqlConfig) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		c.Username, c.Password, c.Addr, c.DBName)
}
