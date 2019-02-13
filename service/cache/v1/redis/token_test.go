package cache_test

import (
	"testing"
)

func TestUpdateToken(t *testing.T) {
	redisCache := getRedisCache(t)

	playerID := "test_playerID"

	token, err := redisCache.UpdateToken(playerID)
	if err != nil {
		t.Errorf("failed to update token : %v", err)
		return
	}

	t.Logf("token = %s", token)
}
