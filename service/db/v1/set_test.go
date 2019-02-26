package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module/v1"
)

func TestSet(t *testing.T) {
	logger.AddDebugLogger()
	cachedDB := getCachedDB(t)

	document := "player"
	where := make(module.Data)
	where["id"] = 3

	data, err := cachedDB.GetOne(document, where)
	if err != nil {
		t.Errorf("failed to get data from db : %v", err)
		return
	}

	data, err = cachedDB.GetOne(document, where)
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
