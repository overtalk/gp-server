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

func TestMysqlDriver_GetUserByAuthCode(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	authCode := "qinhan"
	user, err := mysqlDriver.GetUserByAuthCode(authCode)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", user)
}
