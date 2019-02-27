package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// Test : 测试函数, 以后删除
func (m *MysqlDriver) Test(id int) (*model.Test, error) {
	db := m.conn.First(&model.Test{}, id)
	if db.Error != nil {
		return nil, db.Error
	}

	t := &model.Test{}
	if err := db.Scan(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}
