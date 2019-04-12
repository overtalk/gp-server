package db_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetMatchesNum(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	num, err := mysqlDriver.GetMatchesNum()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(num)
}

func TestMysqlDriver_AddMatch(t *testing.T) {
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

	match := &model.Match{
		IsPublic:     1,
		Title:        "比赛001",
		Introduction: "测试的比赛",
		StartTime:    time.Now().Unix(),
		EndTime:      time.Now().Unix() + 10000,
	}

	if err := mysqlDriver.AddMatch(paper, match); err != nil {
		t.Error(err)
		return
	}
}

func TestMysqlDriver_GetMatches(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var pageIndex int64 = 1
	var pageNum int64 = 3
	matches, err := mysqlDriver.GetMatches(pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}

	for _, match := range matches {
		t.Log(match)
	}
}

func TestMysqlDriver_GetMatchByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 1
	match, err := mysqlDriver.GetMatchByID(id)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(match)
}

func TestMysqlDriver_GetPaperByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 1
	paper, err := mysqlDriver.GetPaperByID(id)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(paper.Paper)
	for _, problem := range paper.P {
		t.Log(problem.Index, problem)
	}
}

func TestMysqlDriver_EditMatch(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 1
	origin, err := mysqlDriver.GetMatchByID(id)
	if err != nil {
		t.Error(err)
		return
	}

	change := &model.Match{
		Id:           id,
		Introduction: "简介在测试中被修改",
	}
	if err := mysqlDriver.EditMatch(change); err != nil {
		t.Error(err)
		return
	}

	changed, err := mysqlDriver.GetMatchByID(id)
	if err != nil {
		t.Error(err)
		return
	}

	if !assert.NotEqual(t, origin.Introduction, changed.Introduction) {
		t.Error("filed [Introduction] is not changed")
		return
	}

	if !assert.Equal(t, changed.Introduction, change.Introduction) {
		t.Error("filed [Introduction] is not changed to expected value")
		return
	}

	if !assert.Equal(t, origin.Title, changed.Title) {
		t.Error("other filed [Title] is changed")
		return
	}
}

func TestAddSomeMatches(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	d := []int64{1, 2}
	bytes, _ := json.Marshal(d)
	for i := 0; i < 10; i++ {
		p := model.Paper{
			Tags:       string(bytes),
			Difficulty: 1,
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

		match := &model.Match{
			IsPublic:     1,
			Title:        "比赛0" + fmt.Sprintf("%d", i),
			Introduction: "测试的比赛",
			StartTime:    time.Now().Unix(),
			EndTime:      time.Now().Unix() + 10000,
		}

		if err := mysqlDriver.AddMatch(paper, match); err != nil {
			t.Error(err)
			return
		}
	}
}
