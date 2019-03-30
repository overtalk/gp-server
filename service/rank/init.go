package rank

import (
	"sync/atomic"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// Rank : implementation of rank module
type Rank struct {
	db    module.DB
	cache module.Cache

	rankItems atomic.Value
}

// NewRank : constructor for module Rank
func NewRank(dataStorage *module.DataStorage) module.Rank {
	p := &Rank{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
	p.initRanks()
	p.loopRanks()
	return p
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewRank(dataStorage)
	gate.RegisterRoute("/rank", "POST", module.GetRankList)
}

func (k *Rank) initRanks() {
	rankItems, err := k.db.GetRank(int(module.MaxRanksNum))
	if err != nil {
		logger.Sugar.Fatalf("failed to get rank info from db : %v", err)
	}
	k.cache.DelRank()
	for _, item := range rankItems {
		if err := k.cache.SetRank(item); err != nil {
			logger.Sugar.Fatalf("failed to set rank info to cache : %v", err)
		}
	}
	k.rankItems.Store(k.turnRankItemsToProtobuf(rankItems))
}
