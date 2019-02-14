package driver

import (
	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Set is to modify db
// if where = nil, create a new record, or update the record
func (p *MysqlDriver) Set(document string, data model.Data, where model.Data) error {
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
