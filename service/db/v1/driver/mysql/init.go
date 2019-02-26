package driver

import (
	"database/sql"
	"errors"
	"fmt"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/QHasaki/gp-server/logger"
)

// MysqlDBInfo describes the details of mysql db
type MysqlDBInfo struct {
	MaxConnection int
	Addr          string
	Username      string
	Password      string
	DBName        string
}

// MysqlDriver describes a mysql connection pool
type MysqlDriver struct {
	dsn           string
	maxConnection int // 设置打开数据库的最大连接数。连接数达到了最大连接数。此时的函数调用将会被block，直到有可用的连接才会返回
	conn          *sql.DB
}

// NewMysqlDriver creates a mysql connection pool
func NewMysqlDriver(dbInfo *MysqlDBInfo) (*MysqlDriver, error) {
	if dbInfo.MaxConnection < 1 {
		logger.Sugar.Errorf("invalid mysql pool size: %d", dbInfo.MaxConnection)
		return nil, errors.New("invalid mysql pool size")
	}

	mysqlDriver := &MysqlDriver{
		maxConnection: dbInfo.MaxConnection,
		dsn: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
			dbInfo.Username, dbInfo.Password, dbInfo.Addr, dbInfo.DBName),
	}

	conn, err := mysqlDriver.CreateDBConn()
	if err != nil {
		logger.Sugar.Errorf("failed to connect to db : %v", err)
		return nil, err
	}

	mysqlDriver.conn = conn

	return mysqlDriver, nil
}

// CreateDBConn creates a new mysql db connection
func (p *MysqlDriver) CreateDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", p.dsn)
	if err != nil {
		logger.Sugar.Errorf("failed to open mysql: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logger.Sugar.Errorf("failed to ping mysql: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(p.maxConnection)

	return db, nil
}

// Ping describes the method to ping db
func (p *MysqlDriver) Ping() error {
	return p.conn.Ping()
}
