package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAllLanguage : get all languag
func (m *MysqlDriver) GetAllLanguage() ([]*model.Language, error) {
	language := make([]*model.Language, 0)
	if err := m.conn.Find(&language); err != nil {
		return nil, err
	}
	return language, nil
}

// AddLanguage : add language
func (m *MysqlDriver) AddLanguage(language *model.Language) error {
	i, err := m.conn.Insert(language)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
