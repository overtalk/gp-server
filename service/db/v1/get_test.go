package db_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
)

func TestGetOne(t *testing.T) {
	cachedDB := getCachedDB(t)

	document := "player"
	columns := []string{"id", "nickname"}
	where := make(model.Data)
	where["id"] = 1

	for i := 0; i < 10; i++ {
		data, err := cachedDB.GetOne(document, columns, where)
		if err != nil {
			t.Errorf("failed to get data from db : %v", err)
			return
		}
		t.Log(data)
	}
}
