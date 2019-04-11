package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_AddDifficulty(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	difficulty := &model.Difficulty{
		Detail: "xxx",
	}

	if err := mysqlDriver.AddDifficulty(difficulty); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", difficulty)
}

func TestAddSomeDifficulty(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	var d = []string{
		"简单",
		"中等",
		"困难",
	}

	for _, v := range d {
		if err := mysqlDriver.AddDifficulty(
			&model.Difficulty{
				Detail: v,
			}); err != nil {
			t.Error(err)
			return
		}
	}
}
