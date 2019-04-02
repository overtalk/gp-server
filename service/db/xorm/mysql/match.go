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
func (m *MysqlDriver) AddMatch(paper *transform.Paper, match *model.Match) error {
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
	// insert match
	match.PaperId = p.Id
	_, err = session.Insert(match)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
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
