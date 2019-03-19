package db

import (
	"strings"

	"github.com/qinhan-shu/gp-server/logger"
)

// Get : get data
// column : at least one column
func (m *MysqlDriver) Get(document string, column []string, where string, args ...interface{}) (map[string]interface{}, error) {
	if !getCheck(column, where, args) {
		return nil, ErrInvaildGetArgs
	}

	// get sql
	sql := "SELECT " + strings.Join(column, ",") + " FROM `" + document + "`"
	if len(where) != 0 {
		sql += " WHERE " + where
	}
	sql += " LIMIT 1"

	// query
	results, err := m.Query(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return nil, err
	}

	return results[0], nil
}

// Gets : get more than one record
func (m *MysqlDriver) Gets(document string, column []string, where string, args ...interface{}) ([]map[string]interface{}, error) {
	if !getCheck(column, where, args) {
		return nil, ErrInvaildGetArgs
	}

	// get sql
	sql := "SELECT " + strings.Join(column, ",") + " FROM `" + document + "`"
	if len(where) != 0 {
		sql += " WHERE " + where
	}

	// query
	results, err := m.Query(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return nil, err
	}

	return results, nil
}

// getCheck : is the input is available
func getCheck(column []string, where string, args ...interface{}) bool {
	if len(column) == 0 {
		return false
	}
	return true
}
