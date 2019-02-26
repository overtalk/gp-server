package driver

import (
	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module/v1"
)

// Set is to modify db
// if where = nil, create a new record, or update the record
func (p *MysqlDriver) Set(document string, data module.Data, where module.Data) error {
	sql, args, err := GetExecSQL(document, data, where)
	if err != nil {
		logger.Sugar.Errorf("failed to get exec sql : %v", err)
		return nil
	}

	resp, err := p.Query(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return err
	}

	defer resp.Close()
	return nil
}
