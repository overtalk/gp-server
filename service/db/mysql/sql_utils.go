package db

import (
	"github.com/qinhan-shu/gp-server/logger"
)

// Exec : update/insert
func (m *MysqlDriver) Exec(sql string, args ...interface{}) error {
	res, err := m.conn.Exec(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to update db: %v", err)
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		logger.Sugar.Errorf("failed to count affected rows: %v", err)
		return err
	}
	if count == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

// Query : select
func (m *MysqlDriver) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := m.conn.Query(query, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query rows: %v", err)
		return nil, err
	}

	var results []map[string]interface{}
	for rows.Next() {
		rowColumns, err := rows.Columns()
		rowColumnCount := len(rowColumns)
		scanFrom := make([]interface{}, rowColumnCount)
		scanTo := make([]interface{}, rowColumnCount)

		for i, _ := range scanFrom {
			scanFrom[i] = &scanTo[i]
		}

		err = rows.Scan(scanFrom...)
		if err != nil {
			return nil, err
		}

		assoc := make(map[string]interface{})

		// Build the associative map from values and column names
		for i, _ := range scanTo {
			assoc[rowColumns[i]] = scanTo[i]
		}

		results = append(results, assoc)
	}
	return results, nil
}
