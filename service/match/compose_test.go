package match_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/match"
)

func TestIntelligentCompose(t *testing.T) {
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	problems, err := dataStorage.DB.GetAllProblems()
	if err != nil {
		t.Error(err)
		return
	}

	paper := &protocol.Paper{
		Difficulty:      1,
		ProblemNum:      3,
		KnowledgePoints: []int64{1, 2},
	}

	paperProblems, err := module.IntelligentCompose(problems, paper)
	if err != nil {
		t.Error(err)
		return
	}

	for _, paperProblem := range paperProblems {
		t.Log(paperProblem)
	}
}
