package db_test

import (
	"fmt"
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetAllTag(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	tags, err := mysqlDriver.GetAllTag()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(tags))
	for _, tag := range tags {
		t.Logf("%+v\n", tag)
	}
}
func TestMysqlDriver_AddTag(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	tag := &model.Tag{
		Detail: "xxx",
	}

	if err := mysqlDriver.AddTag(tag); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", tag)
}

func TestMysqlDriver_UpdateTag(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	tag := &model.Tag{
		Id:     1,
		Detail: "xxx",
	}

	if err := mysqlDriver.UpdateTag(tag); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", tag)
}

func TestAddSomeTags(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 10; i++ {
		if err := mysqlDriver.AddTag(
			&model.Tag{
				Detail: fmt.Sprintf("tag%d", i),
			}); err != nil {
			t.Error(err)
			return
		}
	}
}
