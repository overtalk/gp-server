package paper

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Paper : implementation of paper module
type Paper struct {
	db    module.DB
	cache module.Cache
}

// NewPaper : constructor for paper match
func NewPaper(dataStorage *module.DataStorage) module.Paper {
	return &Paper{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewPaper(dataStorage)
	gate.RegisterRoute("/newPaper", "POST", module.NewPaper)
	gate.RegisterRoute("/modifyPaper", "POST", module.ModifyPaper)
}
