package db_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_AddPaper(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	p := model.Paper{
		ProblemNum:     3,
		Difficulty:     1,
		KnowledgePoint: "并没有填写",
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

	match := &model.Match{
		IsPublic:     1,
		Title:        "比赛001",
		Introduction: "测试的比赛",
		StartTime:    time.Now().Unix(),
		Duration:     10000,
	}

	if err := mysqlDriver.AddMatch(paper, match); err != nil {
		t.Error(err)
		return
	}
}
