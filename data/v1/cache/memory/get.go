package cache

import (
	"github.com/QHasaki/Server/data/v1/error"
	"github.com/QHasaki/Server/utils/serializer"
)

func (c *Cache) GetAll(key string) (map[string]interface{}, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *CacheData:
			var value map[string]interface{}
			if err := serializer.Decode(data.(*CacheData).Data, value); err != nil {
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
