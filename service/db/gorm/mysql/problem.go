package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblems : get problems
func (m *MysqlDriver) GetProblems(tag string) ([]*model.Problem, error) {
	var problems []*model.Problem
	var err error
	if tag == "" {
		err = m.conn.Find(&problems).Error
	} else {
		arg := `%"` + tag + `"%`
		err = m.conn.Where("tags LIKE ?", arg).Find(&problems).Error
	}
	if err != nil {
		return nil, err
	}
	return problems, nil
}

// GetProblemByID : get problem by id
func (m *MysqlDriver) GetProblemByID(id int64) (*model.Problem, error) {
	var p model.Problem
	if err := m.conn.First(&p, id).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *model.Problem) error {
	if checkDefaultValue(problem) && m.conn.NewRecord(problem) {
		return m.conn.Create(problem).Error
	}
	return ErrMissingDefaultValue
}

// UpdateProblem : update problem
func (m *MysqlDriver) UpdateProblem(problem *model.Problem) error {
	return m.conn.Model(problem).Updates(problem).Error
}
