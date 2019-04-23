package problem

import (
	"flag"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

var (
	uploadPath = flag.String("uploadPath", "./tmp", "upload path")
)

// Problem : implementation of Problem module
type Problem struct {
	judgeFilePath string
	db            module.DB
	cache         module.Cache
	path          string
}

// NewProblem : constructor for module Problem
func NewProblem(dataStorage *module.DataStorage) module.Problem {
	logger.Sugar.Infof("file path : %s", *uploadPath)
	return &Problem{
		// judgeFilePath: dataStorage.JudgeFilePath,
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
		path:  *uploadPath,
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
