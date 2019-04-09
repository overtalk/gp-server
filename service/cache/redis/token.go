package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/utils/mode"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

var (
	// expiration defines the expiration time for token
	expiration = 15 * time.Minute
)

// UpdateToken : update token for user
func (r *RedisCache) UpdateToken(userID int64) (string, error) {
	if err := r.DelTokenByUserID(userID); err != nil {
		return "", err
	}

	token := GetToken(userID)
	getTokenKey := getTokenKey(userID)
	getUserIDKey := getUserIDKey(token)

	if _, err := r.client.Set(getTokenKey, token, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of user [%s]", userID)
		return "", err
	}

	if _, err := r.client.Set(getUserIDKey, userID, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of user [%s]", userID)
		return "", err
	}

	return token, nil
}

// GetUserIDByToken : get userID by token
func (r *RedisCache) GetUserIDByToken(token string) (int64, error) {
	// FIXME: remote next line
	return 11, nil
	// for test mode, token is user id itself in test mode
	if mode.GetMode() == mode.TestMode {
		return parse.IntWithError(token)
	}

	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[GetUserIDByToken error] failed to ping redis: %v", err)
		return 0, err
	}

	userIDStr, err := r.client.Get(getUserIDKey(token)).Result()
	if err != nil {
		return 0, err
	}

	userID, err := parse.IntWithError(userIDStr)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// DelTokenByUserID : delete expired token
func (r *RedisCache) DelTokenByUserID(userID int64) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[DelToken error] failed to ping redis: %v", err)
		return err
	}

	// del old token key
	getTokenKey := getTokenKey(userID)
	token, err := r.client.Get(getTokenKey).Result()
	if err == nil {
		// old token is not deleted
		// del the old token
		_, err := r.client.Del(getUserIDKey(token), getTokenKey).Result()
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
	getUserIDKey := getUserIDKey(token)
	userID, err := r.client.Get(getUserIDKey).Result()
	if err == nil {
		// old token is not deleted
		// del the old token
		_, err := r.client.Del(getTokenKey(userID), getUserIDKey).Result()
		if err != nil {
			logger.Sugar.Errorf("[DelToken error] failed to del old token: %v", err)
			return err
		}
	}

	return nil
}
