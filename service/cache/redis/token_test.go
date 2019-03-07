package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTokenAndGetUserIDByToken(t *testing.T) {
	redisCache := getRedisCache(t)

	userID := 1

	token, err := redisCache.UpdateToken(userID)
	if err != nil {
		t.Errorf("failed to update token : %v", err)
		return
	}

	t.Logf("token = %s", token)

	redisUserID, err := redisCache.GetUserIDByToken(token)
	if err != nil {
		t.Errorf("failed to get userID by token : %v", err)
		return
	}

	if !assert.Equal(t, redisUserID, userID) {
		t.Errorf("userID in redis [%d] is not equal to original userID [%d]", redisUserID, userID)
		return
	}
}
