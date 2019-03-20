package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetUserByID : get uer model by user id
func (m *MysqlDriver) GetUserByID(id int64) (*model.User, error) {
	user := new(model.User)
	ok, err := m.conn.Id(id).Get(user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}
	return user, nil
}

// CheckPlayer : get uer model by username and password
func (m *MysqlDriver) CheckPlayer(username, password string) (*model.User, error) {
	user := new(model.User)
	ok, err := m.conn.Alias("u").Where("u.account = ? and u.password = ?", username, password).Get(user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}
	return user, nil
}

// GetUsersByRole : get users by required role
func (m *MysqlDriver) GetUsersByRole(role int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	var err error
	if role < 0 {
		// role < 0 : get all users
		err = m.conn.Find(&users)
	} else {
		err = m.conn.Where("role = ? ", role).Find(&users)
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

// AddUser : add new record
func (m *MysqlDriver) AddUser(user *model.User) error {
	i, err := m.conn.Insert(user)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// UpdateUser : update user
func (m *MysqlDriver) UpdateUser(user *model.User) error {
	affected, err := m.conn.Id(user.Id).Update(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// DelUser : delete user
func (m *MysqlDriver) DelUser(userID int64) error {
	affected, err := m.conn.Id(userID).Delete(&model.User{})
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
