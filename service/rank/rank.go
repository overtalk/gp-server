package rank

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
)

func (k *Rank) loopRanks() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Sugar.Errorf("failed to loopRank: %v", err)
			}
		}()

		var (
			updateTicker = time.NewTicker(module.DefaultRankUpdateInterval)
			cleanTicker  = time.NewTicker(module.DefaultRankCleanInterval)
		)
		for {
			select {
			case <-cleanTicker.C:
				k.cache.CleanRank()
			case <-updateTicker.C:
				k.updateRankItems()
			default:
				time.Sleep(2 * time.Second)
			}
		}
	}()
}

func (k *Rank) updateRankItems() {
	items, err := k.cache.GetRank()
	if err != nil {
		return
	}
	if len(items) > 0 {
		k.rankItems.Store(k.turnRankItemsToProtobuf(items))
	}
}

func (k *Rank) getRanksFromCache() []*protocol.RankItem {
	val, _ := k.rankItems.Load().([]*protocol.RankItem)
	return val
}
