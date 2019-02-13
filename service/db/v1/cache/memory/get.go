package cache

import (
	"github.com/QHasaki/Server/data/v1/error"
	"github.com/QHasaki/Server/utils/serializer"
)

// GetAll is to get cached data from the memery cache
func (c *MemoryCache) GetAll(key string) (map[string]interface{}, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *DataDetails:
			var value map[string]interface{}
			if err := serializer.Decode(data.(*DataDetails).Data, value); err != nil {
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
