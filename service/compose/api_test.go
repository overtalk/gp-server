package compose_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/compose"
	"github.com/qinhan-shu/gp-server/service/config"
)

func TestIntelligentCompose(t *testing.T) {
	db, err := config.NewConfig().GetTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	dataStorage := &module.DataStorage{
		DB: db,
	}
	module := compose.NewCompose(dataStorage)

	problems, err := dataStorage.DB.GetAllProblems()
	if err != nil {
		t.Error(err)
		return
	}

	paper := &protocol.Paper{
		Difficulty: 2,
		ProblemNum: 20,
		Tags:       []int64{1},
	}

	paperProblems, err := module.IntelligentCompose(problems, paper)
	if err != nil {
		t.Error(err)
		return
	}

	var diff int
	var d float64
	for i := 0; i < len(paperProblems); i++ {
		diff += problems[paperProblems[i].ProblemId].Difficulty
	}
	d = float64(diff) / float64(len(paperProblems))

	for _, paperProblem := range paperProblems {
		t.Logf("index = %d, problem id = %d", paperProblem.Index, paperProblem.ProblemId)
	}
	t.Logf("试卷难度系数为：%f", d)

}

func TestRandomCompose(t *testing.T) {
	db, err := config.NewConfig().GetTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	dataStorage := &module.DataStorage{
		DB: db,
	}
	module := compose.NewCompose(dataStorage)

	problems, err := dataStorage.DB.GetAllProblems()
	if err != nil {
		t.Error(err)
		return
	}

	paper := &protocol.Paper{
		Difficulty: 3,
		ProblemNum: 20,
		Tags:       []int64{1, 3, 6},
	}

	paperProblems, err := module.IntelligentCompose(problems, paper)
	if err != nil {
		t.Error(err)
		return
	}

	var diff int
	var d float64
	for i := 0; i < len(paperProblems); i++ {
		diff += problems[paperProblems[i].ProblemId-1].Difficulty
	}
	d = float64(diff) / float64(len(paperProblems))

	for _, paperProblem := range paperProblems {
		t.Logf("index = %d, problem id = %d，difficulty=%d", paperProblem.Index, paperProblem.ProblemId, problems[paperProblem.ProblemId-1].Difficulty)
	}
	t.Logf("试卷难度系数为：%f", d)

}
