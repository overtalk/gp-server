package cache

import (
	"github.com/QHasaki/Server/data/v1/error"
	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/utils/serializer"
)

// GetCache is to get cached data from the memery cache
func (c *MemoryCache) GetCache(key string) (model.Data, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *DataDetails:
			value := make(model.Data)
			if err := serializer.Decode(data.(*DataDetails).Data, &value); err != nil {
				return nil, err
			}
			return value, nil
		case nil:
			return make(model.Data), nil
		default:
			return nil, dataError.ErrInvalidType
		}
	} else {
		return nil, dataError.ErrNoRowsFound
	}
}
