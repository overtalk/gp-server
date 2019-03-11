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

// CheckPlayer : get uer model by username and password
func (m *MysqlDriver) CheckPlayer(username, password string) (*model.User, error) {
	db := m.conn.Where("username = ? and password = ?", username, password).First(&model.User{})
	if db.Error != nil {
		return nil, db.Error
	}

	user := &model.User{}
	if err := db.Scan(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
