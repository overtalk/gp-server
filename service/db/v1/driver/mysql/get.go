package driver

import (
	"github.com/QHasaki/Server/data/v1/error"
	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Get query data from db
func (p *MysqlDriver) Get(document string, column []string, where model.Data) ([]model.Data, error) {
	sql, args := GetQuerySQL(document, column, where)

	resp, err := p.Query(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return nil, err
	}
	defer resp.Close()

	rows := resp.GetRows()

	rowColumns, err := rows.Columns()
	rowColumnCount := len(rowColumns)
	scanAddr := make([]interface{}, rowColumnCount)
	scanResult := make([]interface{}, rowColumnCount)
	for i, _ := range scanAddr {
		scanAddr[i] = &scanResult[i]
	}

	var datas []model.Data

	for rows.Next() {
		err = rows.Scan(scanAddr...)
		if err != nil {
			return nil, err
		}

		assoc := make(model.Data)

		// Build the associative map from values and column names
		for i, _ := range scanResult {
			assoc[rowColumns[i]] = scanResult[i]
		}

		datas = append(datas, assoc)
	}

	if len(datas) == 0 {
		return nil, dataError.ErrNoRowsFound
	}

	return datas, nil
}

// GetOne query data from db
// at most one record
func (p *MysqlDriver) GetOne(document string, column []string, where model.Data) (model.Data, error) {
	sql, args := GetQuerySQL(document, column, where)
	sql += " LIMIT 1"

	resp, err := p.Query(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return nil, err
	}
	defer resp.Close()

	row := resp.GetRows()

	if row.Next() {
		rowColumns, err := row.Columns()

		rowColumnCount := len(rowColumns)
		scanFrom := make([]interface{}, rowColumnCount)
		scanTo := make([]interface{}, rowColumnCount)

		for i, _ := range scanFrom {
			scanFrom[i] = &scanTo[i]
		}

		err = row.Scan(scanFrom...)
		if err != nil {
			return nil, err
		}

		assoc := make(model.Data)

		// Build the associative map from values and column names
		for i, _ := range scanTo {
			assoc[rowColumns[i]] = scanTo[i]
		}

		return assoc, nil
	}
	return nil, dataError.ErrNoRowsFound
}
