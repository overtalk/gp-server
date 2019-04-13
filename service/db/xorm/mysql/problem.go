package db

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblemsNum : get the num of problems
func (m *MysqlDriver) GetProblemsNum(tag int) (int64, error) {
	if tag == 0 {
		m.conn.Count(&model.Problem{})
	}
	return m.conn.Where("tags like ? || tags like ? || tags like ? || tags like ?",
		"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag)).
		Count(&model.Problem{})
}

// GetProblems : get problems
func (m *MysqlDriver) GetProblems(pageNum, pageIndex int64) ([]*transform.IntactProblem, error) {
	problems := make([]*transform.IntactProblem, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Join("INNER", "user", "user.id = problem.publisher").
		Find(&problems); err != nil {
		return nil, err
	}

	for _, problem := range problems {
		testData, err := m.getTestDataByProblemID(problem.Id)
		if err != nil {
			return nil, err
		}
		problem.InAndOutExample = testData
	}
	return problems, nil
}

// GetProblemsByTagID : get problem by tag id
func (m *MysqlDriver) GetProblemsByTagID(pageNum, pageIndex int64, tag int, keyword string) ([]*transform.IntactProblem, error) {
	problems := make([]*transform.IntactProblem, 0)
	if tag == 0 && keyword == "" {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Join("INNER", "user", "user.id = problem.publisher").
			Find(&problems); err != nil {
			return nil, err
		}
	} else if tag != 0 && keyword == "" {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Join("INNER", "user", "user.id = problem.publisher").
			Where("tags like ? || tags like ? || tags like ? || tags like ?",
				"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag)).
			Find(&problems); err != nil {
			return nil, err
		}
	} else if tag == 0 && keyword != "" {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Join("INNER", "user", "user.id = problem.publisher").
			Where("title like ?", "%"+keyword+"%").
			Find(&problems); err != nil {
			return nil, err
		}
	} else {
		if err := m.conn.
			Limit(int(pageNum), int((pageIndex-1)*pageNum)).
			Join("INNER", "user", "user.id = problem.publisher").
			Where("(tags like ? || tags like ? || tags like ? || tags like ?) && title like ?",
				"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag), "%"+keyword+"%").
			Find(&problems); err != nil {
			return nil, err
		}
	}

	for _, problem := range problems {
		testData, err := m.getTestDataByProblemID(problem.Id)
		if err != nil {
			return nil, err
		}
		problem.InAndOutExample = testData
	}
	return problems, nil
}

// AddProblem : add problem to db
func (m *MysqlDriver) AddProblem(problem *transform.IntactProblem) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Insert(&problem.Problem)
	if err != nil {
		session.Rollback()
		return err
	}
	// set problem id
	for _, example := range problem.InAndOutExample {
		example.ProblemId = problem.Id
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
	_, err = session.Id(problem.Id).Update(&problem.Problem)
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
	problem := new(transform.IntactProblem)
	ok, err := m.conn.Id(id).Join("INNER", "user", "user.id = problem.publisher").Get(problem)
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
	problem.InAndOutExample = testData
	return problem, nil
}

// // GetProblemByProblem : get problem by problem
// func (m *MysqlDriver) GetProblemByProblem(problem *model.Problem) (*transform.IntactProblem, error) {
// 	testData, err := m.getTestDataByProblemID(problem.Id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &transform.IntactProblem{
// 		Problem:         *problem,
// 		InAndOutExample: testData,
// 	}, nil
// }

// GetAllProblems : get all
func (m *MysqlDriver) GetAllProblems() ([]*model.Problem, error) {
	problems := make([]*model.Problem, 0)
	if err := m.conn.Cols("id", "tags", "submit_time", "ac").
		Find(&problems); err != nil {
		return nil, err
	}
	return problems, nil
}
