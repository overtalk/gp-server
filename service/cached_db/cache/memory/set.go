package cache

import (
	"bytes"
	"encoding/gob"
	"time"
)

func (c *Cache) SetAll(key string, value map[string]interface{}) error {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(value)
	if err != nil {
		return err
	}
	c.storage.Store(key, &CacheData{
		Data:   network.Bytes(),
		Expire: time.Now().Add(time.Hour),
	})
	return nil
}
