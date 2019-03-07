package utils

import (
	"strings"
)

// RedisKey : create redis key
func RedisKey(args ...string) string {
	return strings.Join(args, ":")
}
