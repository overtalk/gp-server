package cache_test

import (
	"testing"

	"github.com/QHasaki/Server/service/db/v1/cache/memory"
)

func TestSetAllAndGetAll(t *testing.T) {
	memoryCache := cache.NewDBCache()

	key := "DBCache_testKey1"
	value := make(map[string]interface{})
	value["id"] = "test_player_id"
	value["username"] = "test_player_username"

	if err := memoryCache.SetAll(key, value); err != nil {
		t.Errorf("failed to SetAll : %v", err)
		return
	}

	all, err := memoryCache.GetAll(key)
	if err != nil {
		t.Errorf("failed to GetAll : %v", err)
		return
	}

	t.Logf("all cached key : %v", all)
}
