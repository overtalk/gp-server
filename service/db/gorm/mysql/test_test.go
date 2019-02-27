package db_test

import (
	"testing"
)

func TestMysqlDriver_Test(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	test, err := mysqlDriver.Test(1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", test)
}
