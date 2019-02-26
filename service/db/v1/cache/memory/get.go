package cache

import (
	"bytes"
	"encoding/gob"

	"github.com/QHasaki/gp-server/module/v1"
)

// GetCache is to get cached data from the memery cache
func (c *MemoryCache) GetCache(key string) (module.Data, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *CachedData:
			var value map[string]interface{}
			var network bytes.Buffer
			network.Write(data.(*CachedData).Data)
			enc := gob.NewDecoder(&network)
			if err := enc.Decode(&value); err != nil {
				return nil, err
			}
			return value, nil
		case nil:
			return make(module.Data), nil
		default:
			return nil, ErrInvalidType
		}
	} else {
		return nil, ErrNoRowsFound
	}
}
