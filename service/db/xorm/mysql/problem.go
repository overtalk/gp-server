package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblems : get problems
func (m *MysqlDriver) GetProblems(tag string) ([]*model.Problem, error) {
	problems := make([]*model.Problem, 0)
	var err error
	if len(tag) == 0 {
		err = m.conn.Find(&problems)
	} else {
		arg := `%"` + tag + `"%`
		err = m.conn.Where("tags LIKE ?", arg).Find(&problems)
	}
	if err != nil {
		return nil, err
	}
	return problems, nil
}

// GetProblemByID : get problem by id
func (m *MysqlDriver) GetProblemByID(id int64) (*model.Problem, error) {
	problem := new(model.Problem)
	ok, err := m.conn.Id(id).Get(problem)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}
	return problem, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *model.Problem) error {
	i, err := m.conn.Insert(problem)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// UpdateProblem : update problem
func (m *MysqlDriver) UpdateProblem(problem *model.Problem) error {
	affected, err := m.conn.Id(problem.Id).Update(problem)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
