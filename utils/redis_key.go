package utils

import (
	"fmt"
	"strings"
)

// RedisKey : create redis key
func RedisKey(args ...interface{}) string {
	s := make([]string, len(args))
	for i, v := range args {
		s[i] = fmt.Sprint(v)
	}
	return strings.Join(s, ":")
}
