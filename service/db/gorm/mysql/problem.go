package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetProblems : get problems
func (m *MysqlDriver) GetProblems() ([]*model.Problem, error) {
	return nil, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *model.Problem) error {
	if checkDefaultValue(problem) && m.conn.NewRecord(problem) {
		return m.conn.Create(problem).Error
	}
	return ErrMissingDefaultValue
}
