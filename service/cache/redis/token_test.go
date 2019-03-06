package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTokenAndGetPlayerIDByToken(t *testing.T) {
	redisCache := getRedisCache(t)

	playerID := "test_playerID"

	token, err := redisCache.UpdateToken(playerID)
	if err != nil {
		t.Errorf("failed to update token : %v", err)
		return
	}

	t.Logf("token = %s", token)

	redisPlayerID, err := redisCache.GetPlayerIDByToken(token)
	if err != nil {
		t.Errorf("failed to get playerID by token : %v", err)
		return
	}

	if !assert.Equal(t, redisPlayerID, playerID) {
		t.Errorf("playerID in redis [%s] is not equal to original playerID [%s]", redisPlayerID, playerID)
		return
	}
}
