package cache

import (
	"fmt"
)

// getTokenKey is to get to get key for the pair (playerID-token)
// the key is to used to make sure the uniqueness pair for playerID and token
func (r *RedisCache) getPlayerIDToTokenKey(playerID string) string {
	return fmt.Sprintf("player_t_%s", playerID)
}

// getPlayerTokenKey is to get key for the pair (token-playerID)
// the key is to used to authenticate the token, and get playerID
func (r *RedisCache) getTokenToPlayerIDKey(token string) string {
	return fmt.Sprintf("token_p_%s", token)
}
