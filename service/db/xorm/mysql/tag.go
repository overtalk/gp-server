package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAllTag : get all tags
func (m *MysqlDriver) GetAllTag() ([]*model.Tag, error) {
	tags := make([]*model.Tag, 0)
	if err := m.conn.Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// AddTag : add tag
func (m *MysqlDriver) AddTag(tag *model.Tag) error {
	i, err := m.conn.Insert(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
