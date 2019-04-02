package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
)

// AddPaper : add new paper
func (m *MysqlDriver) AddPaper(paper *transform.Paper) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Insert(paper.Paper)
	if err != nil {
		session.Rollback()
		return err
	}
	// set problem id
	for _, problem := range paper.P {
		problem.PaperId = paper.Paper.Id
	}
	_, err = session.Insert(paper.P)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}
