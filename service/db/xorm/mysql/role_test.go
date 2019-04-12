package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetAllRole(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	roles, err := mysqlDriver.GetAllRole()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(roles))
	for _, d := range roles {
		t.Logf("%+v\n", d)
	}
}

func TestAddSomeRole(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	var d = []string{
		"学生",
		"老师",
		"管理员",
	}

	for _, v := range d {
		if err := mysqlDriver.AddRole(
			&model.Role{
				Detail: v,
			}); err != nil {
			t.Error(err)
			return
		}
	}
}
