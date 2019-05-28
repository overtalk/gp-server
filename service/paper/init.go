package paper

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Paper : implementation of paper module
type Paper struct {
	db      module.DB
	cache   module.Cache
	compose module.Compose
}

// NewPaper : constructor for paper match
func NewPaper(dataStorage *module.DataStorage, compose module.Compose) module.Paper {
	return &Paper{
		db:      dataStorage.DB,
		cache:   dataStorage.Cache,
		compose: compose,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage, compose module.Compose) {
	module := NewPaper(dataStorage, compose)
	gate.RegisterRoute("/newPaper", "POST", module.NewPaper)
	gate.RegisterRoute("/modifyPaper", "POST", module.ModifyPaper)
}
