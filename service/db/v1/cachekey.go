package db

import (
	"strings"
	"sync"

	"github.com/QHasaki/Server/module/v1"
	"github.com/QHasaki/Server/utils/parse"
)

var documentCacheKey sync.Map

func init() {
	documentCacheKey.Store("player", "id")
}

// MakeCacheKey is to get cache key
func MakeCacheKey(document string, where module.Data) (string, error) {
	cacheKeyList := []string{document}

	pkName, ok := documentCacheKey.Load(document)
	if !ok {
		return "", ErrMissPK
	}
	pk := parse.String(pkName)

	v, ok := where[pk]
	if !ok {
		return "", ErrMissPK
	}
	cacheKeyList = append(cacheKeyList, pk+"="+parse.String(v))

	return strings.Join(cacheKeyList, "_"), nil
}
