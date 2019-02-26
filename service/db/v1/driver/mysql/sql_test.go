package driver_test

import (
	"testing"

	"github.com/QHasaki/gp-server/module/v1"
	"github.com/QHasaki/gp-server/service/db/v1/driver/mysql"
)

func TestGetQuerySQL(t *testing.T) {
	document := "player"
	where := make(module.Data)

	where["id"] = 1

	t.Log(driver.GetQuerySQL(document, where))
}

func TestGetExecSQL(t *testing.T) {
	document := "player"
	data := make(module.Data)
	data["id"] = 1

	execSQL, args, err := driver.GetExecSQL(document, data, nil)
	if err != nil {
		t.Errorf("failed to get exec sql : %v", err)
		return
	}
	t.Logf("INSERT ----> execSQL = %s, args = %v", execSQL, args)

	where := make(module.Data)

	execSQL, args, err = driver.GetExecSQL(document, data, where)
	if err != nil {
		t.Errorf("failed to get exec sql : %v", err)
		return
	}
	t.Logf("UPDATE(NO WHERE LIMIT) ----> execSQL = %s, args = %v", execSQL, args)

	where["nickname"] = "aaa"

	execSQL, args, err = driver.GetExecSQL(document, data, where)
	if err != nil {
		t.Errorf("failed to get exec sql : %v", err)
		return
	}
	t.Logf("UPDATE(WITH WHERE LIMIT) ----> execSQL = %s, args = %v", execSQL, args)
}
