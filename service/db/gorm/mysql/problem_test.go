package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/gorm"
)

func TestMysqlDriver_AddProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	problem := &model.Problem{
		Title:          "求和问题",
		Description:    "求两个数的和",
		InDescription:  "输入两个int型数",
		OutDescription: "输出两个数的和",
		Hint:           "无提示",
		Example:        `{"test":"test"}`,
		JudgeFile:      "/usr/local/in.out",
		JudgeLimit:     `{"test":"test"}`,
		Tags:           "数组;树",
	}
	if err := mysqlDriver.AddProblem(problem); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", problem)
}
