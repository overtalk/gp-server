package db_test

import (
	"testing"
)

func TestMysqlDriver_GetRank(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	rankItems, err := mysqlDriver.GetRank(9)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range rankItems {
		t.Log(item)
	}
}

func TestMysqlDriver_GetNameAndSubmitNumByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var userID int64 = 1
	name, total, err := mysqlDriver.GetNameAndSubmitNumByID(userID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(name, total)
}
