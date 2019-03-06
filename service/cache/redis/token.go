package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
)

// UpdateToken : update token for player
func (r *RedisCache) UpdateToken(playerID string) (string, error) {
	if err := r.DelToken(playerID); err != nil {
		return "", err
	}

	// expiration defines the expiration time for token
	const expiration time.Duration = 15 * time.Minute

	token := GetToken(playerID)
	playerIDToTokenKey := r.getPlayerIDToTokenKey(playerID)
	tokenToPlayerKey := r.getTokenToPlayerIDKey(token)

	if _, err := r.client.Set(playerIDToTokenKey, token, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of player [%s]", playerID)
		return "", err
	}

	if _, err := r.client.Set(tokenToPlayerKey, playerID, expiration).Result(); err != nil {
		logger.Sugar.Errorf("failed to update token of player [%s]", playerID)
		return "", err
	}

	return token, nil
}

// DelToken : delete expired token
func (r *RedisCache) DelToken(playerID string) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[DelToken error] failed to ping redis: %v", err)
		return err
	}

	// del old token key
	playerIDToTokenKey := r.getPlayerIDToTokenKey(playerID)
	token, err := r.client.Get(playerIDToTokenKey).Result()
	if err == nil {
		// old token is not deleted
		// del the old token
		_, err := r.client.Del(r.getTokenToPlayerIDKey(token)).Result()
		if err != nil {
			logger.Sugar.Errorf("[UpdateToken error] failed to del old token: %v", err)
			return err
		}
	}

	return nil
}

// GetPlayerIDByToken : get playerID by token
func (r *RedisCache) GetPlayerIDByToken(token string) (string, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[GetPlayerIDByToken error] failed to ping redis: %v", err)
		return "", err
	}

	return r.client.Get(r.getTokenToPlayerIDKey(token)).Result()
}
