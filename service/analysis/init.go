package analysis

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Analysis : implementation of analysis module
type Analysis struct {
	db    module.DB
	cache module.Cache
}

// NewAnalysis : constructor for module analysis
func NewAnalysis(dataStorage *module.DataStorage) module.Analysis {
	return &Analysis{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewAnalysis(dataStorage)
	gate.RegisterRoute("/difficultyAnalysis", "POST", module.DifficultyAnalysis)
	gate.RegisterRoute("/tagsAnalysis", "POST", module.TagsAnalysis)
}
