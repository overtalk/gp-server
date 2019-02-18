package driver_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/utils/parse"
)

func TestGet(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"
	where := make(model.Data)

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
	where := make(model.Data)

	where["id"] = 1

	data, err := mysqlDriver.GetOne(document, where)
	if err != nil {
		t.Errorf("failed to get : %v", err)
		return
	}

	t.Logf("id = %d, nickname = %s ", parse.Int(data["id"]), parse.String(data["nickname"]))
}
