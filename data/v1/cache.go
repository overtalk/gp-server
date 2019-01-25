package data

import (
	"strings"

	"github.com/QHasaki/Server/utils/parse"
)

func (p *DB) CacheRefresh(document string, where Data) error {
	key := p.MakeCacheKey(document, where)
	return p.cache.Delete(key)
}

func (p *DB) MakeCacheKey(document string, where Data) string {
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

func (p *DB) CacheDaemon() {
	go p.cache.Daemon()
}
