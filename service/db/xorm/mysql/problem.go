package db

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblemsNum : get the num of problems
func (m *MysqlDriver) GetProblemsNum() (int64, error) {
	return m.conn.Count(&model.Problem{})
}

// GetProblems : get problems
func (m *MysqlDriver) GetProblems(pageNum, pageIndex int64) ([]*transform.IntactProblem, error) {
	problems := make([]*model.Problem, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&problems); err != nil {
		return nil, err
	}

	var intactProblems []*transform.IntactProblem
	for _, problem := range problems {
		intactProblem, err := m.GetProblemByProblem(problem)
		if err != nil {
			return nil, err
		}
		intactProblems = append(intactProblems, intactProblem)
	}

	return intactProblems, nil
}

// GetProblemsByTagID : get problem by tag id
func (m *MysqlDriver) GetProblemsByTagID(pageNum, pageIndex int64, tag int) ([]*transform.IntactProblem, error) {
	problems := make([]*model.Problem, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Where("tags like ? || tags like ? || tags like ? || tags like ?",
			"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag)).
		Find(&problems); err != nil {
		return nil, err
	}

	var intactProblems []*transform.IntactProblem
	for _, problem := range problems {
		testData, err := m.getTestDataByProblemID(problem.Id)
		if err != nil {
			return nil, err
		}
		intactProblems = append(intactProblems, &transform.IntactProblem{
			Detail:          problem,
			InAndOutExample: testData,
		})
	}
	return intactProblems, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *transform.IntactProblem) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Insert(problem.Detail)
	if err != nil {
		session.Rollback()
		return err
	}
	// set problem id
	for _, example := range problem.InAndOutExample {
		example.ProblemId = problem.Detail.Id
	}
	_, err = session.Insert(problem.InAndOutExample)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// UpdateProblem : update problem
func (m *MysqlDriver) UpdateProblem(problem *transform.IntactProblem) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Id(problem.Detail.Id).Update(problem.Detail)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Insert(problem.InAndOutExample)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// GetProblemByID : get problem by id
func (m *MysqlDriver) GetProblemByID(id int64) (*transform.IntactProblem, error) {
	problem := new(model.Problem)
	ok, err := m.conn.Id(id).Get(problem)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}

	testData, err := m.getTestDataByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	return &transform.IntactProblem{
		Detail:          problem,
		InAndOutExample: testData,
	}, nil
}

// GetProblemByProblem : get problem by problem
func (m *MysqlDriver) GetProblemByProblem(problem *model.Problem) (*transform.IntactProblem, error) {
	testData, err := m.getTestDataByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	return &transform.IntactProblem{
		Detail:          problem,
		InAndOutExample: testData,
	}, nil
}

// GetAllProblems : get all
func (m *MysqlDriver) GetAllProblems() ([]*model.Problem, error) {
	problems := make([]*model.Problem, 0)
	if err := m.conn.Cols("id", "tags", "submit_time", "ac").
		Find(&problems); err != nil {
		return nil, err
	}
	return problems, nil
}
