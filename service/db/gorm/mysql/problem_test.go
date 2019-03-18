package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/gorm"
)

func TestMysqlDriver_GetProblems(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	tag := `求和`
	problems, err := mysqlDriver.GetProblems(tag)
	if err != nil {
		t.Error(err)
		return
	}

	for _, problem := range problems {
		t.Logf("%+v\n", problem)
	}
}
func TestMysqlDriver_GetProblemByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 1
	problem, err := mysqlDriver.GetProblemByID(id)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(problem)
}
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
		Example:        `[{"input":"1 1","output":"2"},{"input":"2 2","output":"4"}]`,
		JudgeFile:      "/usr/local/in.out",
		JudgeLimit:     `{"mem": "7m", "time": "62s"}`,
		Tags:           `["求和", "数组", "树"]`,
	}
	if err := mysqlDriver.AddProblem(problem); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", problem)
}

func TestAddSomeProblems(t *testing.T) {
	num := 10
	for i := 0; i < num; i++ {
		TestMysqlDriver_AddProblem(t)
	}
}
