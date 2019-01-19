package cache

import "time"

func (c *Cache) SetAll(key string, value map[string]interface{}) error {
	c.storage.Store(key, &CacheData{
		Data:   value,
		Expire: time.Now().Add(time.Hour),
	})
	return nil
}
