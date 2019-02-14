package driver_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
)

func TestUpdate(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"

	data := make(model.Data)
	data["nickname"] = "aaa"

	where := make(model.Data)
	where["id"] = 1

	if err := mysqlDriver.Set(document, data, where); err != nil {
		t.Errorf("failed to update : %v", err)
		return
	}
}

func TestInsert(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"

	data := make(model.Data)
	data["id"] = 131499
	data["open_id"] = 131499
	data["nickname"] = "aaa"

	if err := mysqlDriver.Set(document, data, nil); err != nil {
		t.Errorf("failed to insert : %v", err)
		return
	}
}
