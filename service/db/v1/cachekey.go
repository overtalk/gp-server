package db

import (
	"strings"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/utils/parse"
)

// MakeCacheKey is to get cache key
func MakeCacheKey(document string, where model.Data) string {
	cacheKeyList := []string{document + "*"}
	primary := []string{
		"id",
		"player_id",
		"open_id",
		"mail_id",
		"season",
	}
	for _, key := range primary {
		if v, ok := where[key]; ok {
			cacheKeyList = append(cacheKeyList, key+"="+parse.String(v))
		}
	}
	return strings.Join(cacheKeyList, "_")
}
