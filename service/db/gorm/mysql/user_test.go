package db_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/model/gorm"
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

	username := "aaa"
	pwd := "aaa"
	user, err := mysqlDriver.CheckPlayer(username, pwd)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", user)
}

func TestMysqlDriver_GetUsersByRole(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var role int64 = -1
	users, err := mysqlDriver.GetUsersByRole(role)
	if err != nil {
		t.Error(err)
		return
	}

	for _, user := range users {
		t.Logf("%+v\n", user)
	}
}

func TestMysqlDriver_AddUser(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	user := &model.User{
		Username:  "test",
		Password:  "test",
		Name:      "test",
		Sex:       false,
		Email:     "test",
		Academy:   "test",
		Major:     "test",
		Create:    time.Now().Unix(),
		LastLogin: time.Now().Unix(),
	}
	if err := mysqlDriver.AddUser(user); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", user)

}
