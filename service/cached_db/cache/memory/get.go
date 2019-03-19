package cache

import (
	"bytes"
	"encoding/gob"

	"github.com/qinhan-shu/gp-server/service/cached_db/error"
)

func (c *Cache) GetAll(key string) (map[string]interface{}, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *CacheData:
			var value map[string]interface{}
			var network bytes.Buffer
			network.Write(data.(*CacheData).Data)
			enc := gob.NewDecoder(&network)
			if err := enc.Decode(&value); err != nil {
				return nil, err
			}
			return value, nil
		case nil:
			return make(map[string]interface{}), nil
		default:
			return nil, dataError.ErrInvalidType
		}
	} else {
		return nil, dataError.ErrNoRowsFound
	}
}
