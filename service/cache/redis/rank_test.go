package cache_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module"
)

func TestRedisCache_SetPveStarRank(t *testing.T) {
	redisCache := getRedisCache(t)

	if err := redisCache.SetRank(&module.RankItem{
		UserID:  1,
		PassNum: 100,
	}); err != nil {
		t.Errorf("SetRank fail : %v\n", err)
		return
	}

	if err := redisCache.SetRank(&module.RankItem{
		UserID:  2,
		PassNum: 10,
	}); err != nil {
		t.Errorf("SetRank fail : %v\n", err)
		return
	}

	if err := redisCache.SetRank(&module.RankItem{
		UserID:  3,
		PassNum: 1000,
	}); err != nil {
		t.Errorf("SetRank fail : %v\n", err)
		return
	}

	items, err := redisCache.GetRank()
	if err != nil {
		t.Errorf("GetRank fail : %v\n", err)
		return
	}
	for i, v := range items {
		t.Logf("rank %d: value: %v\n", i, v)
	}
}
