package cache

import (
	"time"
)

func (c *Cache) Delete(key interface{}) error {
	c.storage.Store(key, nil)
	c.storage.Delete(key)
	return nil
}

func (c *Cache) Daemon() {
	ticker := time.NewTicker(15 * time.Minute)
	for {
		c.storage.Range(func(key, value interface{}) bool {
			if value.(*CacheData).Expire.Unix() <= time.Now().Unix() {
				c.storage.Delete(key)
			}
			return true
		})
		<-ticker.C
	}
}
