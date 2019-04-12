package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAllRole : get all role
func (m *MysqlDriver) GetAllRole() ([]*model.Role, error) {
	roles := make([]*model.Role, 0)
	if err := m.conn.Find(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}

// AddRole : add role
func (m *MysqlDriver) AddRole(tag *model.Role) error {
	i, err := m.conn.Insert(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
