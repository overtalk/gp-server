package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

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
