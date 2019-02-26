package db

import (
	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module/v1"
)

// Set data
func (c *CachedDB) Set(document string, data module.Data, where module.Data) error {
	if err := checkSetCondition(document, data, where); err != nil {
		return err
	}

	cacheKey, err := MakeCacheKey(document, where)
	if err == nil {
		if err := c.cache.DeleteCache(cacheKey); err != nil {
			logger.Sugar.Errorf("failed to del cache key [%s] : %v", cacheKey, err)
		} else {
			logger.Sugar.Debugf("del cache key : %s", cacheKey)
		}
	}

	return c.source.Set(document, data, where)
}
