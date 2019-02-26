package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// UpdateToken is to update token for player
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

// DelToken is to delete expired token
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

// GetToken is to get token key
func GetToken(playerID string) string {
	// in test mode, token = playerID
	if mode.GetMode() == mode.TestMode {
		return playerID
	}

	// return playerID + parse.String(time.Now().Unix()) + parse.String(rand.Int63()+rand.Int63())
	rand1, _ := utils.RandInt(0, 1000000)
	rand2, _ := utils.RandInt(0, 1000000)
	return playerID + "_" + parse.String(time.Now().Unix()) + "_" + parse.String(int64(rand1)+int64(rand2))
}
