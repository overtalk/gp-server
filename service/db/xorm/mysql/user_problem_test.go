package db_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/utils"
)

func TestMysqlDriver_AddSubmitRecord(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	isPass, _ := utils.RandInt(0, 1)
	user, _ := utils.RandInt(1, 10)

	userProblem := &model.UserProblem{
		Code:             "木有代码哦",
		Ispass:           isPass,
		ProblemId:        int64(user),
		RunningLangurage: isPass,
		RunningMem:       10,
		RunningTime:      1,
		SubmitTime:       int(time.Now().Unix()),
		UserId:           int64(user),
	}

	if err := mysqlDriver.AddSubmitRecord(userProblem); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", userProblem)
}

func TestAddSomeUserProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 300; i++ {
		isPass, _ := utils.RandInt(0, 1)
		user, _ := utils.RandInt(1, 10)

		userProblem := &model.UserProblem{
			Code:             "木有代码哦",
			Ispass:           isPass,
			ProblemId:        int64(user),
			RunningLangurage: isPass,
			RunningMem:       10,
			RunningTime:      1,
			SubmitTime:       int(time.Now().Unix()),
			UserId:           int64(user),
		}

		if err := mysqlDriver.AddSubmitRecord(userProblem); err != nil {
			t.Error(err)
			return
		}
	}
}
