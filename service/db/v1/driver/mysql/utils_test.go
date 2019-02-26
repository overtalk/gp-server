package driver_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils/parse"
)

func TestQuery(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	sql := "select `id`,`nickname` from `player` where `id` < ?"
	gender := "10"

	queryResp, err := mysqlDriver.Query(sql, gender)
	if err != nil {
		t.Errorf("failed to Query : %v", err)
		return
	}

	rows := queryResp.GetRows()
	defer queryResp.Close()

	for rows.Next() {
		rowColumns, err := rows.Columns()
		if err != nil {
			t.Errorf("failed to get Columns : %v", err)
			return
		}

		rowColumnCount := len(rowColumns)
		scanAddr := make([]interface{}, rowColumnCount)
		scanResult := make([]interface{}, rowColumnCount)

		for i, _ := range scanAddr {
			scanAddr[i] = &scanResult[i]
		}

		if err = rows.Scan(scanAddr...); err != nil {
			t.Errorf("failed to scan Columns : %v", err)
			return
		}

		t.Logf("id = %d, nickname = %s", parse.Int(scanResult[0]), parse.String(scanResult[1]))
	}
}

func TestQueryRow(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	sql := "select `id`,`nickname` from `player` where `id` < ?"
	gender := "10"

	row := mysqlDriver.QueryRow(sql, gender)

	rowColumnCount := 2
	scanAddr := make([]interface{}, rowColumnCount)
	scanResult := make([]interface{}, rowColumnCount)

	for i, _ := range scanAddr {
		scanAddr[i] = &scanResult[i]
	}

	if err := row.Scan(scanAddr...); err != nil {
		t.Errorf("failed to scan Columns : %v", err)
		return
	}

	t.Logf("id = %d, nickname = %s", parse.Int(scanResult[0]), parse.String(scanResult[1]))

}

func TestExec(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	filed := "CUP_SEASON"
	var season int64

	querySql := "select `value` from `config` where `name` = ?"
	row := mysqlDriver.QueryRow(querySql, filed)
	if err := row.Scan(&season); err != nil {
		t.Errorf("failed to scan : %v", err)
		return
	}

	execSql := "UPDATE `config` SET `value` = ? WHERE `name` = ?"
	args := []interface{}{season + 1, filed}

	if err := mysqlDriver.Exec(execSql, args...); err != nil {
		t.Errorf("failed to exec : %v", err)
		return
	}

	row = mysqlDriver.QueryRow(querySql, filed)
	if err := row.Scan(&season); err != nil {
		t.Errorf("failed to scan : %v", err)
		return
	}

	if !assert.Equal(t, season, parse.Int(args[0])) {
		t.Error("failed to exec")
		return
	}
}
