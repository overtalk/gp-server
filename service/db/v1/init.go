package db

import (
	"github.com/qinhan-shu/gp-server/module/v1"
	"github.com/qinhan-shu/gp-server/service/db/v1/cache/memory"
	"github.com/qinhan-shu/gp-server/service/db/v1/driver/mysql"
)

// CachedDB defines data storage for all the service
// db write operate db directly
// db read search cache(memory / redis ...) first, if not, read db
type CachedDB struct {
	source module.DBDriver
	cache  module.DBCache
}

// NewCachedDB returns CachedDB
func NewCachedDB(mySQLinfo *driver.MysqlDBInfo) (*CachedDB, error) {
	mysqlDriver, err := driver.NewMysqlDriver(mySQLinfo)
	if err != nil {
		return nil, err
	}

	db := &CachedDB{
		cache:  cache.NewDBCache(),
		source: mysqlDriver,
	}

	return db, nil
}

// NoCache return DBDriver, operate db directly
func (c *CachedDB) NoCache() module.DBDriver {
	return c.source
}
