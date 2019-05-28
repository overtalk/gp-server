package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// AddPaper : add new paper
func (m *MysqlDriver) AddPaper(paper *transform.Paper) error {
	// return nil
	p := &paper.Paper
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Insert(p)
	if err != nil {
		session.Rollback()
		return err
	}
	// set problem id
	for _, problem := range paper.P {
		problem.PaperId = p.Id
	}
	_, err = session.Insert(paper.P)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}

// AddPaperProblem : add problem to paper
func (m *MysqlDriver) AddPaperProblem(paperID, problemID int64) error {
	paper, err := m.GetPaperByID(paperID)
	if err != nil {
		return err
	}

	isInclude := false
	maxIndex := 0
	for _, v := range paper.P {
		if v.Index > maxIndex {
			maxIndex = v.Index
		}
		if v.ProblemId == problemID {
			isInclude = true
		}
	}

	if isInclude {
		return nil
	}

	newProblem := &model.PaperProblem{
		Index:     maxIndex + 1,
		PaperId:   paperID,
		ProblemId: problemID,
	}

	session := m.conn.NewSession()
	defer session.Close()

	err = session.Begin()
	if err != nil {
		session.Rollback()
		return err
	}
	// 增加paper——problem
	_, err = session.Insert(newProblem)
	if err != nil {
		session.Rollback()
		return err
	}
	// 修改题目数量
	_, err = session.Id(paper.Id).Update(&model.Paper{ProblemNum: paper.ProblemNum + 1})
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// DelPaperProblem : del problem from paper
func (m *MysqlDriver) DelPaperProblem(paperID, problemID int64) error {
	paper, err := m.GetPaperByID(paperID)
	if err != nil {
		return err
	}

	isInclude := false
	maxIndex := 0
	for _, v := range paper.P {
		if v.Index > maxIndex {
			maxIndex = v.Index
		}
		if v.ProblemId == problemID {
			isInclude = true
		}
	}

	if !isInclude {
		return nil
	}

	session := m.conn.NewSession()
	defer session.Close()

	err = session.Begin()
	if err != nil {
		session.Rollback()
		return err
	}
	// 增加paper——problem
	_, err = m.conn.Where("problem_id = ? and paper_id = ?", problemID, paperID).Delete(&model.PaperProblem{})
	if err != nil {
		session.Rollback()
		return err
	}
	// 修改题目数量
	_, err = session.Id(paper.Id).Update(&model.Paper{ProblemNum: paper.ProblemNum - 1})
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}
