package db

import (
	"github.com/QHasaki/Server/model/v1"
)

// Set data
func (c *CachedDB) Set(document string, data model.Data, where model.Data) error {
	if err := setCondition(document, data, where); err != nil {
		return err
	}

	return c.source.Set(document, data, where)
}
