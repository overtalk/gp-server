package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module/v1"
)

func TestGetOne(t *testing.T) {
	logger.AddDebugLogger()
	cachedDB := getCachedDB(t)

	document := "player"
	where := make(module.Data)
	where["id"] = 1

	for i := 0; i < 10; i++ {
		_, err := cachedDB.GetOne(document, where)
		if err != nil {
			t.Errorf("failed to get data from db : %v", err)
			return
		}
		//t.Log(data)
	}
}
