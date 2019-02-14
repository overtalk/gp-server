package cache_test

import (
	"testing"
	"time"

	"github.com/QHasaki/Server/service/db/v1/cache/memory"
)

func TestDeleteCache(t *testing.T) {
	memoryCache := cache.NewDBCache()

	key := "DBCache_testKey1"
	value := make(map[string]interface{})
	value["id"] = "test_player_id"
	value["username"] = "test_player_username"

	if err := memoryCache.SetCache(key, value); err != nil {
		t.Errorf("failed to SetAll : %v", err)
		return
	}

	if err := memoryCache.DeleteCache(key); err != nil {
		t.Errorf("failed to Delete : %v", err)
		return
	}

	_, err := memoryCache.GetCache(key)
	if err == nil {
		t.Errorf("failed to Delete : %v", err)
		return
	}
}

func TestDaemon(t *testing.T) {
	memoryCache := cache.NewDBCache()

	cache.DefaultCacheDuration = 1 * time.Second
	cache.DefaultClearDuration = 2 * time.Second
	go memoryCache.Daemon()

	key := "DBCache_testKey1"
	value := make(map[string]interface{})
	value["id"] = "test_player_id"
	value["username"] = "test_player_username"

	if err := memoryCache.SetCache(key, value); err != nil {
		t.Errorf("failed to SetAll : %v", err)
		return
	}

	time.Sleep(3 * time.Second)

	_, err := memoryCache.GetCache(key)
	if err == nil {
		t.Errorf("failed to delete in Daemon : %v", err)
		return
	}
}
