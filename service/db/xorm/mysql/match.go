package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetMatchesNum : get the num of match
func (m *MysqlDriver) GetMatchesNum() (int64, error) {
	return m.conn.Count(&model.Match{})
}

// AddMatch : add new match
func (m *MysqlDriver) AddMatch(match *model.Match) error {
	// insert match
	_, err := m.conn.Insert(match)
	return err
}

// GetMatches : get matches
func (m *MysqlDriver) GetMatches(pageNum, pageIndex int64) ([]*model.Match, error) {
	matches := make([]*model.Match, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Find(&matches); err != nil {
		return nil, err
	}

	return matches, nil
}

// GetMatchByID : get match by id
func (m *MysqlDriver) GetMatchByID(id int64) (*model.Match, error) {
	match := new(model.Match)
	ok, err := m.conn.Id(id).Get(match)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}

	return match, nil
}

// GetPaperByID : get paper info
func (m *MysqlDriver) GetPaperByID(id int64) (*transform.Paper, error) {
	paper := new(model.Paper)
	ok, err := m.conn.Id(id).Get(paper)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}

	paperProblems := make([]*model.PaperProblem, 0)
	if err := m.conn.Where("paper_id = ?", id).
		Asc("index").
		Find(&paperProblems); err != nil {
		return nil, err
	}

	problems := make([]*transform.IntactProblem, 0)
	for _, paperProblem := range paperProblems {
		problem, err := m.GetProblemByID(paperProblem.ProblemId)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}

	return &transform.Paper{
		Paper:          *paper,
		P:              paperProblems,
		ProblemsDetail: problems,
	}, nil
}

// EditMatch : edit match info
func (m *MysqlDriver) EditMatch(match *model.Match) error {
	affected, err := m.conn.Id(match.Id).Update(match)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
