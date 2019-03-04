package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetMatchByID : get uer model by match id
func (m *MysqlDriver) GetMatchByID(id int) (*model.Match, error) {
	db := m.conn.First(&model.Match{}, id)
	if db.Error != nil {
		return nil, db.Error
	}

	match := &model.Match{}
	if err := db.Scan(match).Error; err != nil {
		return nil, err
	}

	return match, nil
}
