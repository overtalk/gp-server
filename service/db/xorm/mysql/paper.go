package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	// "github.com/qinhan-shu/gp-server/model/xorm"
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
