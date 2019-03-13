package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetMatchByID : get uer model by match id
func (m *MysqlDriver) GetMatchByID(id int) (*model.Match, error) {
	var match model.Match
	if err := m.conn.First(&match, id).Error; err != nil {
		return nil, err
	}
	return &match, nil
}
