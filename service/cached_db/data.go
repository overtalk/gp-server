package data

import (
	"github.com/qinhan-shu/gp-server/service/cached_db/cache/memory"
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
)

type Data = map[string]interface{}
type DataInfo = driver.DBInfo
type CacheInfo = cache.CacheInfo
type Cache = cache.Cache
type Where = driver.Where

func NewConnect(info DataInfo, cacheInfo cache.CacheInfo) (*CachedDB, error) {
	pool, err := driver.Connect(info)
	if err != nil {
		return nil, err
	}
	cachePool, err := cache.Connect(cacheInfo)
	if err != nil {
		return nil, err
	}

	gameData := &CachedDB{
		origin: &originDB{
			conn: pool,
		},
		cache: cachePool,
	}
	return gameData, nil
}

func (p *CachedDB) NoCache() *originDB {
	return p.origin
}
