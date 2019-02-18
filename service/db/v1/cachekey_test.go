package db_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/service/db/v1"
)

func TestMakeCacheKey(t *testing.T) {
	document := "player"
	where := make(model.Data)
	where["id"] = 1
	where["nickname"] = "qinhan"

	t.Log(db.MakeCacheKey(document, where))
}
