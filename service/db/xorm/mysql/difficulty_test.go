package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetAllDifficulty(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	difficulty, err := mysqlDriver.GetAllDifficulty()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(difficulty))
	for _, d := range difficulty {
		t.Logf("%+v\n", d)
	}
}
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
		"很简单",
		"不那么简单",
		"中等",
		"有点困难",
		"很困难",
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
