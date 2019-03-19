package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/model/gorm"
	"github.com/qinhan-shu/gp-server/utils"
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
		Account:   "test",
		Password:  "test",
		Role:      0,
		Name:      "test",
		Sex:       false,
		Email:     "test",
		Phone:     "909",
		School:    "shu",
		Create:    time.Now().Unix(),
		LastLogin: time.Now().Unix(),
	}
	if err := mysqlDriver.AddUser(user); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", user)
}

func TestMysqlDriver_UpdateUser(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var userID int64 = 1
	originUser, err := mysqlDriver.GetUserByID(userID)
	if err != nil {
		t.Error(err)
		return
	}

	change := &model.User{
		ID:   userID,
		Name: originUser.Name + "000",
	}
	if err := mysqlDriver.UpdateUser(change); err != nil {
		t.Error(err)
		return
	}

	changedUser, err := mysqlDriver.GetUserByID(userID)
	if err != nil {
		t.Error(err)
		return
	}

	if !assert.NotEqual(t, originUser.Name, changedUser.Name) {
		t.Error("filed [Name] is not changed")
		return
	}

	if !assert.Equal(t, changedUser.Name, change.Name) {
		t.Error("filed [Name] is not changed to expected value")
		return
	}

	if !assert.Equal(t, originUser.Account, changedUser.Account) {
		t.Error("other filed [Username] is changed")
		return
	}
}

func TestMysqlDriver_DelUser(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	newUser := &model.User{
		Account:   "test12",
		Password:  "test",
		Role:      1,
		Name:      "test",
		Sex:       false,
		Email:     "test",
		Phone:     "test",
		School:    "test",
		Create:    time.Now().Unix(),
		LastLogin: time.Now().Unix(),
	}
	if err := mysqlDriver.AddUser(newUser); err != nil {
		t.Error(err)
		return
	}

	if err := mysqlDriver.DelUser(newUser.ID); err != nil {
		t.Error(err)
		return
	}

	_, err = mysqlDriver.GetUserByID(newUser.ID)
	if err == nil {
		t.Error("failed to delete player")
		return
	}
}

func TestAddSomeUsers(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	keyWord := "jack"
	num := 10
	for i := 0; i < num; i++ {
		key := keyWord + fmt.Sprintf("%d", i)
		role, _ := utils.RandInt(0, 2)
		user := &model.User{
			Account:   key,
			Password:  utils.MD5(key),
			Role:      role,
			Name:      key,
			Sex:       role == 0,
			Phone:     key,
			Email:     key,
			School:    key,
			Create:    time.Now().Unix(),
			LastLogin: time.Now().Unix(),
		}
		if err := mysqlDriver.AddUser(user); err != nil {
			t.Error(err)
			return
		}
	}
}
