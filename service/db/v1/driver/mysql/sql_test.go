package driver_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/service/db/v1/driver/mysql"
)

func TestGetQuerySQL(t *testing.T) {
	document := "player"
	columns := []string{"id", "nickname"}
	where := make(model.Data)

	where["id"] = 1

	t.Log(driver.GetQuerySQL(document, columns, where))
}

func TestGetExecSQL(t *testing.T) {
	document := "player"
	data := make(model.Data)
	data["id"] = 1

	execSQL, args := driver.GetExecSQL(document, data, nil)
	t.Logf("INSERT ----> execSQL = %s, args = %v", execSQL, args)

	where := make(model.Data)

	execSQL, args = driver.GetExecSQL(document, data, where)
	t.Logf("UPDATE(NO WHERE LIMIT) ----> execSQL = %s, args = %v", execSQL, args)

	where["nickname"] = "aaa"

	execSQL, args = driver.GetExecSQL(document, data, where)
	t.Logf("UPDATE(WITH WHERE LIMIT) ----> execSQL = %s, args = %v", execSQL, args)
}
