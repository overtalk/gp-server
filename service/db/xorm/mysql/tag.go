package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// AddTag : add tag
func (m *MysqlDriver) AddTag(tag *model.Tag) error {
	i, err := m.conn.Insert(tag)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

func (m *MysqlDriver) getTagsByProblemID(problemID int64) ([]*model.ProblemTag, error) {
	problemTags := make([]*model.ProblemTag, 0)
	if err := m.conn.Where("problem_id = ?", problemID).Find(&problemTags); err != nil {
		return nil, err
	}
	return problemTags, nil
}
