package db_test

import (
	"encoding/json"
	"testing"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_AddPaper(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	d := []int64{1, 2}
	bytes, _ := json.Marshal(d)

	p := model.Paper{
		Difficulty: 1,
		Tags:       string(bytes),
		ProblemNum: 3,
	}
	paper := &transform.Paper{
		Paper: p,
		P: []*model.PaperProblem{
			&model.PaperProblem{
				Index:     1,
				ProblemId: 1,
			},
			&model.PaperProblem{
				Index:     2,
				ProblemId: 2,
			},
			&model.PaperProblem{
				Index:     3,
				ProblemId: 4,
			},
		},
	}

	if err := mysqlDriver.AddPaper(paper); err != nil {
		t.Error(err)
		return
	}
}

func TestMysqlDriver_AddPaperProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var paperID int64 = 12
	var problemID int64 = 3

	if err := mysqlDriver.AddPaperProblem(paperID, problemID); err != nil {
		t.Error(err)
		return
	}
}

func TestMysqlDriver_DelPAperProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var paperID int64 = 12
	var problemID int64 = 3

	if err := mysqlDriver.DelPaperProblem(paperID, problemID); err != nil {
		t.Error(err)
		return
	}
}
