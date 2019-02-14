package db

import (
	"github.com/QHasaki/Server/model/v1"
)

// Get get data
// TODO: get from DBCache first, if not found, query db source
func (c *CachedDB) Get(document string, column []string, where model.Data) ([]model.Data, error) {
	if err := checkGetCondition(document, column, where); err != nil {
		return nil, err
	}

	return c.source.Get(document, column, where)
}

// GetOne get from DBCache first
// if not founded in DBCache, query from db
func (c *CachedDB) GetOne(document string, column []string, where model.Data) (model.Data, error) {
	if err := checkGetCondition(document, column, where); err != nil {
		return nil, err
	}

	return c.source.GetOne(document, column, where)
}
