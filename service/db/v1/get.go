package db

import (
	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Get get data
// TODO: get from DBCache first, if not found, query db source
func (c *CachedDB) Get(document string, where model.Data) ([]model.Data, error) {
	if err := checkGetCondition(document, where); err != nil {
		return nil, err
	}

	return c.source.Get(document, where)
}

// GetOne get from DBCache first
// if not founded in DBCache, query from db
func (c *CachedDB) GetOne(document string, where model.Data) (model.Data, error) {
	if err := checkGetCondition(document, where); err != nil {
		return nil, err
	}

	isCached := true
	cacheKey, err := MakeCacheKey(document, where)
	if err != nil {
		isCached = false
	}

	if isCached {
		data, err := c.cache.GetCache(cacheKey)
		if err == nil {
			logger.Sugar.Debugf("get key [ %s ] from cache", cacheKey)
			return data, nil
		}
	}

	data, err := c.source.GetOne(document, where)
	if err != nil {
		return nil, err
	}

	if err := c.cache.SetCache(cacheKey, data); err != nil {
		logger.Sugar.Errorf("failed to set cache key [%s] : %v", cacheKey, err)
		return data, nil
	}

	return data, err
}
