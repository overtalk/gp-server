package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/config/v1"
	"github.com/qinhan-shu/gp-server/service/db/gorm/mysql"
)

func TestMysqlDriver_Connect(t *testing.T) {
	configs := config.NewConfig()

	mysqlAddr, err := configs.GetConfigByName("MYSQL_ADDR")
	if err != nil {
		t.Error(err)
		return
	}

	mysqlUser, err := configs.GetConfigByName("MYSQL_USER")
	if err != nil {
		t.Error(err)
		return
	}

	mysqlPassword, err := configs.GetConfigByName("MYSQL_PASS")
	if err != nil {
		t.Error(err)
		return
	}

	mysqlDBName, err := configs.GetConfigByName("MYSQL_DBNAME")
	if err != nil {
		t.Error(err)
		return
	}

	c := &db.MysqlConfig{
		Addr:     mysqlAddr,
		Username: mysqlUser,
		Password: mysqlPassword,
		DBName:   mysqlDBName,
	}

	_, err = db.NewMysqlDriver(c)
	if err != nil {
		t.Error(err)
		return
	}
}
