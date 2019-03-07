package cache

import (
	"fmt"
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

var (
	// expiration defines the expiration time for token
	expiration = 15 * time.Minute
)

// UpdateToken : update token for user
func (r *RedisCache) UpdateToken(userID int) (string, error) {
	if err := r.DelTokenByUserID(userID); err != nil {
		return "", err
	}

	token := GetToken(userID)
	userIDToTokenKey := r.getUserIDToTokenKey(fmt.Sprintf("%d", userID))
	tokenToUserKey := r.getTokenToUserIDKey(token)

	if _, err := r.client.Set(userIDToTokenKey, token, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of user [%s]", userID)
		return "", err
	}

	if _, err := r.client.Set(tokenToUserKey, userID, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of user [%s]", userID)
		return "", err
	}

	return token, nil
}

// GetUserIDByToken : get userID by token
func (r *RedisCache) GetUserIDByToken(token string) (int, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[GetUserIDByToken error] failed to ping redis: %v", err)
		return 0, err
	}

	userIDStr, err := r.client.Get(r.getTokenToUserIDKey(token)).Result()
	if err != nil {
		return 0, err
	}

	userID, err := parse.IntWithError(userIDStr)
	if err != nil {
		return 0, err
	}
	return int(userID), nil
}

// DelTokenByUserID : delete expired token
func (r *RedisCache) DelTokenByUserID(userID int) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[DelToken error] failed to ping redis: %v", err)
		return err
	}

	// del old token key
	userIDToTokenKey := r.getUserIDToTokenKey(fmt.Sprintf("%d", userID))
	token, err := r.client.Get(userIDToTokenKey).Result()
	if err == nil {
		// old token is not deleted
		// del the old token
		_, err := r.client.Del(r.getTokenToUserIDKey(token), userIDToTokenKey).Result()
		if err != nil {
			logger.Sugar.Errorf("[DelToken error] failed to del old token: %v", err)
			return err
		}
	}

	return nil
}

// DelTokenByToken : delete expired token
func (r *RedisCache) DelTokenByToken(token string) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[DelToken error] failed to ping redis: %v", err)
		return err
	}

	// del old token key
	tokenToUserIDKey := r.getTokenToUserIDKey(token)
	userID, err := r.client.Get(tokenToUserIDKey).Result()
	if err == nil {
		// old token is not deleted
		// del the old token
		_, err := r.client.Del(r.getUserIDToTokenKey(userID), tokenToUserIDKey).Result()
		if err != nil {
			logger.Sugar.Errorf("[DelToken error] failed to del old token: %v", err)
			return err
		}
	}

	return nil
}
