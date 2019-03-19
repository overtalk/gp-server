package db

import (
	"strings"
)

// Set : update/insert data
// column : at least one column
func (m *MysqlDriver) Set(document string, data map[string]interface{}, where string, args ...interface{}) error {
	if !setCheck(data) {
		return ErrInvaildGetArgs
	}

	if len(where) == 0 {
		// insert
		var (
			column []string
			needed []string
			values []interface{}
		)

		for k, v := range data {
			column = append(column, "`"+k+"`")
			needed = append(needed, "?")
			values = append(values, v)
		}

		sql := "INSERT INTO `" + document + "` "
		sql += "(" + strings.Join(column, ",") + ") VALUES (" + strings.Join(needed, ", ") + ")"

		return m.Exec(sql, values...)
	}
	// update
	var (
		datas  []string
		values []interface{}
	)
	for k, v := range data {
		datas = append(datas, k+" = ?")
		values = append(values, v)
	}

	sql := "UPDATE `" + document + "` SET "
	sql += strings.Join(datas, ", ") + " WHERE " + where

	return m.Exec(sql, append(values, args...)...)
}

// setCheck : is the input is available
func setCheck(data map[string]interface{}) bool {
	if len(data) == 0 {
		return false
	}
	return true
}
