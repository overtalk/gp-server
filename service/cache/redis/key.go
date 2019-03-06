package cache

import (
	"fmt"
)

// getTokenKey is to get to get key for the pair (userID-token)
// the key is to used to make sure the uniqueness pair for userID and token
func (r *RedisCache) getUserIDToTokenKey(userID string) string {
	return fmt.Sprintf("u_%s", userID)
}

// getTokenToUserIDKey is to get key for the pair (token-userID)
// the key is to used to authenticate the token, and get userID
func (r *RedisCache) getTokenToUserIDKey(token string) string {
	return fmt.Sprintf("t_%s", token)
}
