package db

import (
	model_utils "github.com/qinhan-shu/gp-server/model"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblems : get problems
func (m *MysqlDriver) GetProblems(pageNum, pageIndex int64) ([]*model_utils.IntactProblem, error) {
	problems := make([]*model.Problem, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&problems); err != nil {
		return nil, err
	}

	var intactProblems []*model_utils.IntactProblem
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
func (m *MysqlDriver) GetProblemsByTagID(pageNum, pageIndex int64, tag int) ([]*model_utils.IntactProblem, error) {
	problemTags := make([]*model.ProblemTag, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Where("tag_id = ?", tag).Find(&problemTags); err != nil {
		return nil, err
	}

	var intactProblems []*model_utils.IntactProblem
	for _, tag := range problemTags {
		intactProblem, err := m.GetProblemByID(tag.ProblemId)
		if err != nil {
			return nil, err
		}
		intactProblems = append(intactProblems, intactProblem)
	}

	return intactProblems, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *model_utils.IntactProblem) error {
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
	// set problem id
	for _, tag := range problem.Tags {
		tag.ProblemId = problem.Detail.Id
	}
	_, err = session.Insert(problem.Tags)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// UpdateProblem : update problem
func (m *MysqlDriver) UpdateProblem(problem *model_utils.IntactProblem) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Where("id = ?", problem.Detail.Id).Update(problem.Detail)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Insert(problem.InAndOutExample)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Insert(problem.Tags)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// GetProblemByID : get problem by id
func (m *MysqlDriver) GetProblemByID(id int64) (*model_utils.IntactProblem, error) {
	problem := new(model.Problem)
	ok, err := m.conn.Id(id).Get(problem)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}

	tags, err := m.getTagsByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	testData, err := m.getTestDataByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	return &model_utils.IntactProblem{
		Detail:          problem,
		InAndOutExample: testData,
		Tags:            tags,
	}, nil
}

// GetProblemByProblem : get problem by problem
func (m *MysqlDriver) GetProblemByProblem(problem *model.Problem) (*model_utils.IntactProblem, error) {
	tags, err := m.getTagsByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	testData, err := m.getTestDataByProblemID(problem.Id)
	if err != nil {
		return nil, err
	}
	return &model_utils.IntactProblem{
		Detail:          problem,
		InAndOutExample: testData,
		Tags:            tags,
	}, nil
}
