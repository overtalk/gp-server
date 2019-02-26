package driver_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module/v1"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

func TestGet(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"
	where := make(module.Data)

	datas, err := mysqlDriver.Get(document, where)
	if err != nil {
		t.Errorf("failed to get : %v", err)
		return
	}

	for _, data := range datas {
		t.Logf("id = %d, nickname = %s ", parse.Int(data["id"]), parse.String(data["nickname"]))
	}
}

func TestGetOne(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"
	where := make(module.Data)

	where["id"] = 1

	data, err := mysqlDriver.GetOne(document, where)
	if err != nil {
		t.Errorf("failed to get : %v", err)
		return
	}

	t.Logf("id = %d, nickname = %s ", parse.Int(data["id"]), parse.String(data["nickname"]))
}
