package db

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetProblemsNum : get the num of problems
func (m *MysqlDriver) GetProblemsNum(tag int, keyword string) (int64, error) {
	if tag == 0 && keyword == "" {
		return m.conn.Count(&model.Problem{})
	} else if tag == 0 && keyword != "" {
		return m.conn.Where("title like ?", "%"+keyword+"%").Count(&model.Problem{})
	} else if tag != 0 && keyword == ""  {
		return m.conn.Where("tags like ? || tags like ? || tags like ? || tags like ?",
			"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag)).
			Count(&model.Problem{})
	}
	return m.conn.Where("(tags like ? || tags like ? || tags like ? || tags like ?) && title like ?",
		"%"+fmt.Sprintf(",%d,", tag)+"%", fmt.Sprintf("[%d,", tag)+"%", "%"+fmt.Sprintf(",%d]", tag), fmt.Sprintf("[%d]", tag), "%"+keyword+"%").
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

	nextIndex, err := m.getNextTestDataIndex(problem.Problem.Id)
	if err != nil {
		session.Rollback()
		return err
	}
	// set problem id
	for _, example := range problem.InAndOutExample {
		nextIndex++
		example.Index = nextIndex
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
	affected, err := session.Id(problem.Id).Update(&problem.Problem)
	fmt.Println("affected rows in update problem : ", affected)
	if err != nil {
		session.Rollback()
		return err
	}

	nextIndex, err := m.getNextTestDataIndex(problem.Problem.Id)
	if err != nil {
		session.Rollback()
		return err
	}
	// get insert and update test data
	var (
		toInsert []*model.TestData
		toUpdate []*model.TestData
	)
	for _, v := range problem.InAndOutExample {
		if v.Index == 0 {
			nextIndex++
			v.Index = nextIndex
			toUpdate = append(toUpdate, v)
		} else {
			toInsert = append(toInsert, v)
		}
	}
	_, err = session.Insert(toInsert)
	if err != nil {
		session.Rollback()
		return err
	}

	_, err = session.Update(toUpdate)
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
	if err := m.conn.Cols("id", "tags", "submit_time", "ac", "difficulty", "last_used").
		Find(&problems); err != nil {
		return nil, err
	}
	return problems, nil
}

// GetNextTestDataIndex : get max index
func (m *MysqlDriver) getNextTestDataIndex(id int64) (int64, error) {
	return m.conn.Where("problem_id = ?", id).Count(&model.TestData{})
}
