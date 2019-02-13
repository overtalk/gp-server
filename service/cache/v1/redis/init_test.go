package cache_test

import (
	"testing"

	"github.com/QHasaki/Server/service/cache/v1/redis"
	"github.com/QHasaki/Server/service/config/v1"
)

func TestPing(t *testing.T) {
	c := config.NewConfig()

	addr, err := c.GetConfigByConfigName("REDIS_ADDR")
	if err != nil {
		t.Error(err)
		return
	}
	pass, err := c.GetConfigByConfigName("REDIS_PASS")
	if err != nil {
		t.Error(err)
		return
	}

	redisCache, err := cache.NewRedisCache(addr, pass, 1)
	if err != nil {
		t.Errorf("failed to new redis cache : %v", err)
		return
	}

	if err := redisCache.Ping(); err != nil {
		t.Errorf("failed to ping redis cache : %v", err)
		return
	}
}
