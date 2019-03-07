package cache

import (
	"github.com/qinhan-shu/gp-server/utils"
)

// getTokenKey is the key to get token by userID
func getTokenKey(userID interface{}) string {
	return utils.RedisKey("TOKEN", "getToken", userID)
}

// getUserIDKey is the key to get userID by token
func getUserIDKey(token string) string {
	return utils.RedisKey("TOKEN", "getUserID", token)
}
