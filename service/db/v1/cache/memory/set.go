package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/qinhan-shu/gp-server/module/v1"
)

// DefaultCacheDuration describes the default
var DefaultCacheDuration = time.Hour

// SetCache is to set data to memory
func (c *MemoryCache) SetCache(key string, value module.Data) error {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(value)
	if err != nil {
		return err
	}
	c.storage.Store(key, &CachedData{
		Data:   network.Bytes(),
		Expire: time.Now().Add(time.Hour),
	})
	return nil
}
