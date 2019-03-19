package manage

import (
	"github.com/qinhan-shu/gp-server/module"
)

// BackStageManage : implementation of backstage manager module
type BackStageManage struct {
	judgeFilePath string
	db            module.DB
	cache         module.Cache
}

// NewBackStageManager : constructor for module BackStageManager
func NewBackStageManager(dataStorage *module.DataStorage) module.BackStageManage {
	return &BackStageManage{
		judgeFilePath: dataStorage.JudgeFilePath,
		db:            dataStorage.DB,
		cache:         dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewBackStageManager(dataStorage)
	// user manage
	gate.RegisterRoute("/getUsers", module.GetUsers)
	gate.RegisterRoute("/addUsers", module.AddUsers)
	gate.RegisterRoute("/updateUsers", module.UpdateUsers)
	gate.RegisterRoute("/delUsers", module.DelUsers)
	// problem manage
	gate.RegisterRoute("/getProblems", module.GetProblems)
	gate.RegisterRoute("/getProblemByID", module.GetProblemByID)
	gate.RegisterRoute("/addProblem", module.AddProblem)
	gate.RegisterRoute("/editProblem", module.EditProblem)
}
