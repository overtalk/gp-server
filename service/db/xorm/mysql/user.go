package db

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetUsersNum : get the num of users
func (m *MysqlDriver) GetUsersNum(role int64) (int64, error) {
	if role == 0 {
		return m.conn.Count(&model.User{})
	}
	return m.conn.Where("role = ?", role).Count(&model.User{})
}

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

// GetUsers : get users
func (m *MysqlDriver) GetUsers(pageNum, pageIndex int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByRole : get users by required role
func (m *MysqlDriver) GetUsersByRole(pageNum, pageIndex, role int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := m.conn.
		Where("role = ?", role).
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&users); err != nil {
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

// GetSubmitRecord : get recrod from db
func (m *MysqlDriver) GetSubmitRecord(userID, problemID, pageNum, pageIndex int64) ([]*model.UserProblem, int64, error) {
	var (
		num int64
		err error
	)
	records := make([]*model.UserProblem, 0)

	if userID != 0 && problemID == 0 {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Where("user_id = ?", userID).
			Find(&records); err != nil {
			return nil, 0, err
		}
		num, err = m.conn.Where("user_id = ?", userID).Count(&model.UserProblem{})
	} else if userID == 0 && problemID != 0 {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Where("problem_id = ?", problemID).
			Find(&records); err != nil {
			return nil, 0, err
		}
		num, err = m.conn.Where("problem_id = ?", problemID).Count(&model.UserProblem{})
	} else if userID != 0 && problemID != 0 {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Where("problem_id = ? and user_id = ?", problemID, userID).
			Find(&records); err != nil {
			return nil, 0, err
		}
		num, err = m.conn.Where("problem_id = ? and user_id = ?", problemID, userID).Count(&model.UserProblem{})
	} else {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Find(&records); err != nil {
			return nil, 0, err
		}
		num, err = m.conn.Count(&model.UserProblem{})
	}
	if err != nil {
		logger.Sugar.Errorf("failed to get the count of user_problem : %v", err)
		return nil, 0, err
	}

	return records, num, nil
}

// CreatePlayer : create a new player
func (m *MysqlDriver) CreatePlayer(user *model.User) error {
	user.Role = 1
	user.Create = time.Now().Unix()
	user.LastLogin = time.Now().Unix()

	i, err := m.conn.Insert(user)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
