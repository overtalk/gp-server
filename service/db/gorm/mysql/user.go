package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetUserByID : get uer model by user id
func (m *MysqlDriver) GetUserByID(id int) (*model.User, error) {
	db := m.conn.First(&model.User{}, id)
	if db.Error != nil {
		return nil, db.Error
	}

	user := &model.User{}
	if err := db.Scan(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
