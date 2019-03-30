package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetRank : get rank info from db
func (m *MysqlDriver) GetRank(num int) ([]*module.RankItem, error) {
	sql := "SELECT  user_id , count(*)  FROM user_problem where `isPass` = 1 group by `user_id` ORDER BY count(*) DESC limit ?;"
	results, err := m.conn.QueryInterface(sql, num)
	if err != nil {
		return nil, err
	}
	var rankItems []*module.RankItem
	for _, v := range results {
		rankItems = append(rankItems, &module.RankItem{
			UserID:  parse.Int(v["user_id"]),
			PassNum: parse.Int(v["count(*)"]),
		})
	}
	return rankItems, nil
}

// GetNameAndSubmitNumByID : get name and submit num
func (m *MysqlDriver) GetNameAndSubmitNumByID(userID int64) (string, int64, error) {
	user := new(model.User)
	ok, err := m.conn.Cols("name").Id(userID).Get(user)
	if err != nil {
		return "", 0, err
	}
	if !ok {
		return "", 0, ErrNoRowsFound
	}

	total, err := m.conn.Where("user_id = ?", userID).Count(&model.UserProblem{})
	if err != nil {
		return "", 0, err
	}

	return user.Name, total, nil
}
