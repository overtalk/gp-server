package cache

import (
	"time"

	"github.com/QHasaki/Server/utils/serializer"
)

func (c *Cache) SetAll(key string, value map[string]interface{}) error {
	data, err := serializer.Encode(value)
	if err != nil {
		return err
	}
	c.storage.Store(key, &CacheData{
		Data:   data,
		Expire: time.Now().Add(time.Hour),
	})
	return nil
}
