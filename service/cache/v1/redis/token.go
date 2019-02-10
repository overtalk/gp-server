package cache

import (
	"math/rand"
	"time"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/utils/mode"
	"github.com/QHasaki/Server/utils/parse"
)

// UpdateToken is to update token for player
func (r *RedisCache) UpdateToken(playerID string) (string, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("[UpdateToken error] failed to ping redis: %v", err)
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

	// TODO: refresh memory cache
	// if err := c.dataDB.RefreshByPlayerID(playerID); err != nil {
	// 	sugar.Errorf("refresh player cache failed : %v", err)
	// }

	return token, nil
}

// GetToken is to get token key
func GetToken(playerID string) string {
	// in test mode, token = playerID
	if mode.GetMode() == mode.TestMode {
		return playerID
	}

	return playerID + parse.String(time.Now().Unix()) + parse.String(rand.Int63()+rand.Int63())
}
