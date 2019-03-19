package db

import (
	"strings"
)

// Get : get data
// column : at least one column
func (m *MysqlDriver) Get(document string, column []string, where ...Condition) (map[string]interface{}, error) {
	if !getCheck(column) {
		return nil, ErrInvaildGetArgs
	}

	whereStr, args := TurnMysqlWhere(where)

	// get sql
	sql := "SELECT " + strings.Join(column, ",") + " FROM `" + document + "`"
	if len(whereStr) != 0 {
		sql += " WHERE " + whereStr
	}
	sql += " LIMIT 1"

	// query
	results, err := m.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	return results[0], nil
}

// Gets : get more than one record
func (m *MysqlDriver) Gets(document string, column []string, where ...Condition) ([]map[string]interface{}, error) {
	if !getCheck(column) {
		return nil, ErrInvaildGetArgs
	}

	whereStr, args := TurnMysqlWhere(where)

	// get sql
	sql := "SELECT " + strings.Join(column, ",") + " FROM `" + document + "`"
	if len(where) != 0 {
		sql += " WHERE " + whereStr
	}

	// query
	results, err := m.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// getCheck : is the input is available
func getCheck(column []string) bool {
	if len(column) == 0 {
		return false
	}
	return true
}
