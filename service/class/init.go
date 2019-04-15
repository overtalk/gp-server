package class

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Class : implementation of backstage class module
type Class struct {
	judgeFilePath string
	db            module.DB
	cache         module.Cache
}

// NewClass : constructor for module Class
func NewClass(dataStorage *module.DataStorage) module.Class {
	return &Class{
		// judgeFilePath: dataStorage.JudgeFilePath,
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewClass(dataStorage)
	gate.RegisterRoute("/getClasses", "POST", module.GetClasses)
	gate.RegisterRoute("/getClassByID", "POST", module.GetClassByID)
	gate.RegisterRoute("/addClass", "POST", module.AddClass)
	gate.RegisterRoute("/editClass", "POST", module.EditClass)
	gate.RegisterRoute("/memberManage", "POST", module.EditClass)
	gate.RegisterRoute("/getMembers", "POST", module.GetMembers)
	gate.RegisterRoute("/enterClass", "POST", module.EnterClass)
}
