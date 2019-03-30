package db

import (
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetRank : get rank info from db
func (m *MysqlDriver) GetRank(num int) ([]module.RankItem, error) {
	sql := "SELECT  user_id , count(*)  FROM user_problem where `isPass` = 1 group by `user_id` ORDER BY count(*) DESC limit ?;"
	results, err := m.conn.QueryInterface(sql, num)
	if err != nil {
		return nil, err
	}
	var rankItems []module.RankItem
	for _, v := range results {
		rankItems = append(rankItems, module.RankItem{
			UserID:  parse.Int(v["user_id"]),
			PassNum: parse.Int(v["count(*)"]),
		})
	}
	return rankItems, nil
}
