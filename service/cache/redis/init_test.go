package cache_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module/v1"
	"github.com/qinhan-shu/gp-server/service/cache/v1/redis"
	"github.com/qinhan-shu/gp-server/service/config/v1"
)

func getRedisCache(t *testing.T) module.Cache {
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
