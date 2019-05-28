package db_test

import (
	"testing"
)

func TestMysqlDriver_GetAlgorithm(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	algorithms, err := mysqlDriver.GetAlgorithm()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(algorithms))
	for _, d := range algorithms {
		t.Logf("%+v\n", d)
	}
}
