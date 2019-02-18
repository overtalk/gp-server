package cache_test

import (
	"testing"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/service/cache/v1/redis"
	"github.com/QHasaki/Server/service/config/v1"
)

func getRedisCache(t *testing.T) model.Cache {
	c := config.NewConfig()

	addr, err := c.GetConfigByName("REDIS_ADDR")
	if err != nil {
		t.Error(err)
		return nil
	}
	pass, err := c.GetConfigByName("REDIS_PASS")
	if err != nil {
		t.Error(err)
		return nil
	}

	redisCache, err := cache.NewRedisCache(addr, pass, 1)
	if err != nil {
		t.Errorf("failed to new redis cache : %v", err)
		return nil
	}

	return redisCache
}

func TestPing(t *testing.T) {
	redisCache := getRedisCache(t)
	if err := redisCache.Ping(); err != nil {
		t.Errorf("failed to ping redis cache : %v", err)
		return
	}
}
