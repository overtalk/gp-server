package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func (m *MysqlDriver) getTestDataByProblemID(problemID int64) ([]*model.TestData, error) {
	testData := make([]*model.TestData, 0)
	if err := m.conn.Where("problem_id = ?", problemID).Find(&testData); err != nil {
		return nil, err
	}
	return testData, nil
}
