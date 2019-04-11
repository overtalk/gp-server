package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetAllCognition(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	cognition, err := mysqlDriver.GetAllCognition()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(cognition))
	for _, c := range cognition {
		t.Logf("%+v\n", c)
	}
}
func TestMysqlDriver_AddCognition(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	cognition := &model.Cognition{
		Detail: "xxx",
	}

	if err := mysqlDriver.AddCognition(cognition); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", cognition)
}

func TestAddSomeCognition(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	var d = []string{
		"识记",
		"理解",
		"应用",
	}

	for _, v := range d {
		if err := mysqlDriver.AddCognition(
			&model.Cognition{
				Detail: v,
			}); err != nil {
			t.Error(err)
			return
		}
	}
}
