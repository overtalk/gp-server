package driver_test

import (
	"testing"

	"github.com/QHasaki/Server/service/db/v1/driver/mysql"
)

func TestNewMysqlDriver(t *testing.T) {
	dbInfo := &driver.MysqlDBInfo{
		MaxConnection: 2000,
		Addr:          "172.26.32.12",
		Username:      "sausage",
		Password:      "sausage_shooter",
		DBName:        "sausage_shooter",
	}
	mysqlDriver, err := driver.NewMysqlDriver(dbInfo)
	if err != nil {
		t.Errorf("failed to new mysql driver : %v", err)
		return
	}

	if err := mysqlDriver.Ping(); err != nil {
		t.Errorf("failed to ping mysql : %v", err)
		return
	}
}
