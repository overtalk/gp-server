package db_test

import (
	"testing"
)

func TestMysqlDriver_GetUserByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	user, err := mysqlDriver.GetUserByID(1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", user)
}

func TestMysqlDriver_CheckPlayer(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	username := "jack"
	pwd := "jack"
	user, err := mysqlDriver.CheckPlayer(username, pwd)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", user)
}
