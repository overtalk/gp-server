package problem

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Problem : implementation of Problem module
type Problem struct {
	judgeFilePath string
	db            module.DB
	cache         module.Cache
}

// NewProblem : constructor for module Problem
func NewProblem(dataStorage *module.DataStorage) module.Problem {
	return &Problem{
		// judgeFilePath: dataStorage.JudgeFilePath,
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewProblem(dataStorage)

	gate.RegisterRoute("/getProblems", "POST", module.GetProblems)
	gate.RegisterRoute("/getProblemByID", "POST", module.GetProblemByID)
	gate.RegisterRoute("/addProblem", "POST", module.AddProblem)
	gate.RegisterRoute("/editProblem", "POST", module.EditProblem)
}
