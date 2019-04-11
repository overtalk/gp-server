package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAllCognition : get all difficulty
func (m *MysqlDriver) GetAllCognition() ([]*model.Cognition, error) {
	cognition := make([]*model.Cognition, 0)
	if err := m.conn.Find(&cognition); err != nil {
		return nil, err
	}
	return cognition, nil
}

// AddCognition : add cognition
func (m *MysqlDriver) AddCognition(tag *model.Cognition) error {
	i, err := m.conn.Insert(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
