package db_test

import (
	"testing"

	"github.com/QHasaki/Server/service/db/v1"
	"github.com/QHasaki/Server/service/db/v1/driver/mysql"
)

func getCachedDB(t *testing.T) *db.CachedDB {
	dbInfo := &driver.MysqlDBInfo{
		MaxConnection: 2000,
		Addr:          "localhost",
		Username:      "root",
		Password:      "12345678",
		DBName:        "sausage_shooter",
	}

	cachedDB, err := db.NewCachedDB(dbInfo)
	if err != nil {
		t.Errorf("failed to new cached db : %v ", err)
		return nil
	}

	return cachedDB
}

func TestNewCachedDB(t *testing.T) {
	getCachedDB(t)
}
