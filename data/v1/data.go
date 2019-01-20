package data

import (
	"github.com/QHasaki/Server/data/v1/cache/memory"
	"github.com/QHasaki/Server/data/v1/driver/mysql"
)

type Data = map[string]interface{}
type DataInfo = driver.DBInfo
type CacheInfo = cache.CacheInfo
type Cache = cache.Cache
type Where = driver.Where

func NewConnect(info DataInfo, cacheInfo cache.CacheInfo) (*DB, error) {
	pool, err := driver.Connect(info)
	if err != nil {
		return nil, err
	}
	cachePool, err := cache.Connect(cacheInfo)
	if err != nil {
		return nil, err
	}

	gameData := &DB{
		origin: &originDB{
			conn: pool,
		},
		cache: cachePool,
	}
	return gameData, nil
}

func (p *DB) NoCache() *originDB {
	return p.origin
}
