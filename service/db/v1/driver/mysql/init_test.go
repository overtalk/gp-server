package driver_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/db/v1/driver/mysql"
)

func getMySqlDriver(t *testing.T) *driver.MysqlDriver {
	dbInfo := &driver.MysqlDBInfo{
		MaxConnection: 2000,
		Addr:          "localhost",
		Username:      "root",
		Password:      "12345678",
		DBName:        "sausage_shooter",
	}
	mysqlDriver, err := driver.NewMysqlDriver(dbInfo)
	if err != nil {
		t.Errorf("failed to new mysql driver : %v", err)
		return nil
	}
	return mysqlDriver
}

func TestNewMysqlDriver(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)
	if err := mysqlDriver.Ping(); err != nil {
		t.Errorf("failed to ping mysql : %v", err)
		return
	}
}
