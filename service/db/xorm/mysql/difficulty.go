package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAllDifficulty : get all difficulty
func (m *MysqlDriver) GetAllDifficulty() ([]*model.Difficulty, error) {
	difficulty := make([]*model.Difficulty, 0)
	if err := m.conn.Find(&difficulty); err != nil {
		return nil, err
	}
	return difficulty, nil
}

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
