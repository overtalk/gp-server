package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func (m *MysqlDriver) getTagsByProblemID(problemID int64) ([]*model.ProblemTag, error) {
	problemTags := make([]*model.ProblemTag, 0)
	if err := m.conn.Where("problem_id = ?", problemID).Find(&problemTags); err != nil {
		return nil, err
	}
	return problemTags, nil
}

