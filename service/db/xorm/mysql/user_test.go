package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/utils"
)

func TestMysqlDriver_GetUsersNum(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	num, err := mysqlDriver.GetUsersNum(2)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("the num of users : %d", num)
}
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

	username := "jack1"
	pwd := "jack1"
	user, err := mysqlDriver.CheckPlayer(username, utils.MD5(pwd))
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

	var role int64 = 1
	var pageIndex int64 = 1
	var pageNum int64 = 3
	users, err := mysqlDriver.GetUsersByRole(pageNum, pageIndex, role)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(users))
	for _, user := range users {
		t.Logf("%+v\n", user)
	}
}

func TestMysqlDriver_GetUsers(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var pageIndex int64 = 1
	var pageNum int64 = 3
	users, err := mysqlDriver.GetUsers(pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(users))
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
		Account:   "lyx",
		Password:  utils.MD5("lyx"),
		Role:      0,
		Name:      "李亿璇",
		Sex:       0,
		Email:     "xxx@xxx.com",
		Phone:     "xxxxxxxxxx",
		School:    "shu",
		Create:    (time.Now().Unix()),
		LastLogin: (time.Now().Unix()),
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
		Id:   (userID),
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
		Sex:       0,
		Email:     "test",
		Phone:     "test",
		School:    "test",
		Create:    (time.Now().Unix()),
		LastLogin: (time.Now().Unix()),
	}
	if err := mysqlDriver.AddUser(newUser); err != nil {
		t.Error(err)
		return
	}

	if err := mysqlDriver.DelUser(newUser.Id); err != nil {
		t.Error(err)
		return
	}

	_, err = mysqlDriver.GetUserByID(newUser.Id)
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

	keyWord := "tom"
	num := 10
	for i := 0; i < num; i++ {
		key := keyWord + fmt.Sprintf("%d", i)
		role, _ := utils.RandInt(1, 3)
		sex, _ := utils.RandInt(0, 1)
		user := &model.User{
			Account:   key,
			Password:  utils.MD5(key),
			Role:      role,
			Name:      key,
			Sex:       sex,
			Phone:     key,
			Email:     key,
			School:    key,
			Create:    time.Now().Unix(),
			LastLogin: time.Now().Unix(),
		}
		if i == 0 {
			user.Role = 3
		}
		if err := mysqlDriver.AddUser(user); err != nil {
			t.Error(err)
			return
		}
	}

	sex, _ := utils.RandInt(0, 1)
	user := &model.User{
		Account:   "1",
		Password:  utils.MD5("1"),
		Role:      2,
		Name:      "lyx",
		Sex:       sex,
		Phone:     "111",
		Email:     "xxx@xxx",
		School:    "shu",
		Create:    time.Now().Unix(),
		LastLogin: time.Now().Unix(),
	}
	if err := mysqlDriver.AddUser(user); err != nil {
		t.Error(err)
		return
	}
}

func TestMysqlDriver_GetSubmitRecord(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var userID int64 = 1
	var problemID int64 = 1
	var pageIndex int64 = 1
	var pageNum int64 = 3
	records, num, err := mysqlDriver.GetSubmitRecord(userID, problemID, pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("the num of  all records : %d", num)
	t.Logf("the num of records : %d", len(records))

	for _, record := range records {
		t.Log(record.Id)
		t.Log(record.UserId)
		t.Log(record.ProblemId)
		t.Log(record)
	}
}

func TestMysqlDriver_CreatePlayer(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	user := &model.User{
		Account:  "qqq",
		Password: "qqq",
		Name:     "qinhan",
		
	}

	if err := mysqlDriver.CreatePlayer(user); err != nil {
		t.Error(err)
		return
	}
}
