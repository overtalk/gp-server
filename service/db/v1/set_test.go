package db_test

import (
	"testing"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

func TestSet(t *testing.T) {
	logger.AddDebugLogger()
	cachedDB := getCachedDB(t)

	document := "player"
	columns := []string{"id", "nickname"}
	where := make(model.Data)
	where["id"] = 3

	data, err := cachedDB.GetOne(document, columns, where)
	if err != nil {
		t.Errorf("failed to get data from db : %v", err)
		return
	}

	data, err = cachedDB.GetOne(document, columns, where)
	if err != nil {
		t.Errorf("failed to get data from db : %v", err)
		return
	}

	data["nickname"] = "test"

	if err := cachedDB.Set(document, data, where); err != nil {
		t.Errorf("failed to set data to db : %v", err)
		return
	}
}
