package conf

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Conf : implementation of conf module
type Conf struct {
	db    module.DB
	cache module.Cache
}

// NewConf : constructor for module conf
func NewConf(dataStorage *module.DataStorage) module.Conf {
	return &Conf{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewConf(dataStorage)
	gate.RegisterRoute("/conf", "GET", module.GetConfig)
	gate.RegisterRoute("/userRole", "GET", module.GetUserRole)
	gate.RegisterRoute("/getAllLanguage", "GET", module.GetAllLanguage)
	gate.RegisterRoute("/getJudgeResult", "GET", module.GetJudgeResult)
	gate.RegisterRoute("/paperComposeAlgorithm", "GET", module.GetAlgorithm)
}
