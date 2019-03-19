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
	gate.RegisterRoute("/getUsers", "POST", module.GetUsers)
	gate.RegisterRoute("/addUsers", "POST", module.AddUsers)
	gate.RegisterRoute("/updateUsers", "POST", module.UpdateUsers)
	gate.RegisterRoute("/delUsers", "POST", module.DelUsers)
	// problem manage
	gate.RegisterRoute("/getProblems", "POST", module.GetProblems)
	gate.RegisterRoute("/getProblemByID", "POST", module.GetProblemByID)
	gate.RegisterRoute("/addProblem", "POST", module.AddProblem)
	gate.RegisterRoute("/editProblem", "POST", module.EditProblem)
}
