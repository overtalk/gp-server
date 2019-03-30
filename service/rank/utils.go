package rank

import (
	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
)

func (k *Rank) turnRankItemsToProtobuf(items []*module.RankItem) []*protocol.RankItem {
	var rankItems []*protocol.RankItem
	for index, item := range items {
		name, submitNum, err := k.db.GetNameAndSubmitNumByID(item.UserID)
		if err != nil {
			logger.Sugar.Errorf("failed to get name and submit num for user %d", item.UserID)
			continue
		}
		rankItems = append(rankItems, &protocol.RankItem{
			Ranking:   int64(index) + 1,
			UserID:    item.UserID,
			PassNum:   item.PassNum,
			Name:      name,
			SubmitNum: submitNum,
		})
	}
	return rankItems
}
