package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetAlgorithm : get all algorithm
func (m *MysqlDriver) GetAlgorithm() ([]*model.Algorithm, error) {
	algorithms := make([]*model.Algorithm, 0)
	if err := m.conn.Find(&algorithms); err != nil {
		return nil, err
	}
	return algorithms, nil
}
