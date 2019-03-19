package data

import (
	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
)

// CachedDB : db with cache
type CachedDB struct {
	origin *originDB
	cache  *Cache
}

// Set : set data
func (p *CachedDB) Set(document string, data Data, where Data) error {
	if err := getCondition(document, data, where); err != nil {
		return err
	}
	cacheKey := p.MakeCacheKey(document, where)
	if err := driver.Set(p.origin.conn, document, data, where); err != nil {
		return err
	}
	if err := p.cache.Delete(cacheKey); err != nil {
		logger.Sugar.Errorf("cant delete data cache : %v", err)
	}
	return nil
}

// Get : get data
func (p *CachedDB) Get(document string, column []string, where Data) (Data, error) {
	if err := setCondition(document, column, where); err != nil {
		return nil, err
	}
	cacheKey := p.MakeCacheKey(document, where)
	cacheData, err := p.cache.GetAll(cacheKey)
	if err != nil || len(cacheData) < 1 {
		data, err := driver.Get(p.origin.conn, document, []string{"*"}, where)
		if err != nil {
			return nil, err
		}
		err = p.cache.SetAll(cacheKey, data)
		if err != nil {
			logger.Sugar.Errorf("cant save to data cache: %v", err)
		}
		return data, nil
	}
	return cacheData, nil
}

// Inc : increase certaion field
func (p *CachedDB) Inc(document string, column []string, where Data) error {
	if err := driver.Inc(p.origin.conn, document, column, where); err != nil {
		return err
	}
	return nil
}

// GetAll : get all record
func (p *CachedDB) GetAll(document string, column []string, where Data) ([]Data, error) {
	return driver.GetAll(p.origin.conn, document, column, where)
}
