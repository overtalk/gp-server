package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/db/xorm/mysql"
)

func TestMysqlDriver_Connect(t *testing.T) {
	_, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
}

func getMysqlDriver() (module.DB, error) {
	configs := config.NewConfig()

	keys := []string{
		"MYSQL_ADDR",
		"MYSQL_USER",
		"MYSQL_PASS",
		"MYSQL_DBNAME",
	}

	serverCfg := make(map[string]string)

	for _, key := range keys {
		value, err := configs.GetConfigByName(key)
		if err != nil {
			return nil, err
		}
		serverCfg[key] = value
	}

	c := &db.MysqlConfig{
		Addr:     serverCfg["MYSQL_ADDR"],
		Username: serverCfg["MYSQL_USER"],
		Password: serverCfg["MYSQL_PASS"],
		DBName:   serverCfg["MYSQL_DBNAME"],
	}

	return db.NewMysqlDriver(c)
}
