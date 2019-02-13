package cache

import (
	"time"

	"github.com/QHasaki/Server/utils/serializer"
)

// DefaultCacheDuration describes the default
var DefaultCacheDuration = time.Hour

// SetAll is to set data to memory
func (c *MemoryCache) SetAll(key string, value map[string]interface{}) error {
	data, err := serializer.Encode(value)
	if err != nil {
		return err
	}
	c.storage.Store(key, &DataDetails{
		Data:   data,
		Expire: time.Now().Add(DefaultCacheDuration),
	})
	return nil
}
