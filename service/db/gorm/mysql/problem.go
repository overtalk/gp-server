package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetProblems : get problems
func (m *MysqlDriver) GetProblems() ([]*model.Problem, error) {
	var problems []*model.Problem
	if err := m.conn.Find(&problems).Error; err != nil {
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
