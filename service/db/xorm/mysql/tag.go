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

// GetAllTagsByPage : get all tags by page
func (m *MysqlDriver) GetAllTagsByPage(pageNum, pageIndex int64) ([]*model.Tag, error) {
	tags := make([]*model.Tag, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&tags); err != nil {
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

// UpdateTag : update tag
func (m *MysqlDriver) UpdateTag(tag *model.Tag) error {
	i, err := m.conn.Id(tag.Id).Update(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
