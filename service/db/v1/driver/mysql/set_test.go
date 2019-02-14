package driver_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
)

func TestSet(t *testing.T) {
	mysqlDriver := getMySqlDriver(t)

	document := "player"

	data := make(model.Data)
	data["nickname"] = "qqq"

	where := make(model.Data)
	where["id"] = 1

	if err := mysqlDriver.Set(document, data, where); err != nil {
		t.Errorf("failed to set : %v", err)
		return
	}

}
