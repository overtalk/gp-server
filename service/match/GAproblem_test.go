package match_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/match"
)

func TestInit(t *testing.T) {
	db, err := config.NewConfig().GetTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	dataStorage := &module.DataStorage{
		DB: db,
	}

	problems, err := dataStorage.DB.GetAllProblems()
	if err != nil {
		t.Error(err)
		return
	}
	 problem := match.GAproblem{}
	problem.Init(1, problems)
	t.Log(problem)

}
