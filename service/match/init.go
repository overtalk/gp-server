package match

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Match : implementation of match module
type Match struct {
	db    module.DB
	cache module.Cache
}

// NewMatch : constructor for module match
func NewMatch(dataStorage *module.DataStorage) module.Match {
	return &Match{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewMatch(dataStorage)
	gate.RegisterRoute("/newMatch", "POST", module.NewMatch)
	gate.RegisterRoute("/editMatch", "POST", module.EditMatch)
	gate.RegisterRoute("/getMatchByID", "POST", module.GetMatchByID)
	gate.RegisterRoute("/getPaperByID", "POST", module.GetPaperByID)
	gate.RegisterRoute("/getMatches", "POST", module.GetMatches)
}
