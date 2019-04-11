package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// AddDifficulty : add difficulty
func (m *MysqlDriver) AddDifficulty(tag *model.Difficulty) error {
	i, err := m.conn.Insert(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
