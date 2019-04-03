package db

import (
	"fmt"
)

var (
	// DefaultMaxOpenConnsNum : The default maximum number of mysql connections
	DefaultMaxOpenConnsNum = 20
	// DefaultMaxIdleConnsNum : The default maximum number of mysql idle connections
	DefaultMaxIdleConnsNum = 0
)

// MysqlConfig : mysql database configuration
type MysqlConfig struct {
	MaxOpenConnsNum int // maximum open connections
	MaxIdleConnsNum int // maximum number of idle connections
	Addr            string
	Username        string
	Password        string
	DBName          string
	IsCached        bool // is cache
}

func (c *MysqlConfig) getDSN() string {
	fmt.Printf("%s:%s@/%s?charset=utf8",
		c.Username, c.Password, c.DBName)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		c.Username, c.Password, c.Addr, c.DBName)
}
