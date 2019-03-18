package manage

import (
	"github.com/qinhan-shu/gp-server/module"
)

// BackStageManage : implementation of backstage manager module
type BackStageManage struct {
	db    module.DB
	cache module.Cache
}

// NewBackStageManager : constructor for module BackStageManager
func NewBackStageManager(dataStorage *module.DataStorage) module.BackStageManage {
	return &BackStageManage{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	managerModule := NewBackStageManager(dataStorage)
	// user manage
	gate.RegisterRoute("/getUsers", module.Router{
		Method:  "POST",
		Handler: managerModule.GetUsers,
	})
	gate.RegisterRoute("/addUsers", module.Router{
		Method:  "POST",
		Handler: managerModule.AddUsers,
	})
	gate.RegisterRoute("/updateUsers", module.Router{
		Method:  "POST",
		Handler: managerModule.UpdateUsers,
	})
	gate.RegisterRoute("/delUsers", module.Router{
		Method:  "POST",
		Handler: managerModule.DelUsers,
	})
}
