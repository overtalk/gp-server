package driver

import (
	"strings"

	"github.com/QHasaki/Server/data/v1/error"
	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Get query data from db
func (p *MysqlDriver) Get(document string, column []string, where model.Data) (model.Data, error) {
	var (
		columns string
		wheres  []string
		values  []interface{}
	)
	sql := "SELECT "

	columns = strings.Join(column, ",")

	for k, v := range where {
		switch v.(type) {
		case model.Where:
			wheres = append(wheres, k+" "+v.(model.Where).Operator+" ?")
			values = append(values, v.(model.Where).Value)
		default:
			wheres = append(wheres, k+" = ?")
			values = append(values, v)
		}
	}

	resp, err := p.Query(sql+columns+" FROM `"+document+"` WHERE "+strings.Join(wheres, " AND "), values...)
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
