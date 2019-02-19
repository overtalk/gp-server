package cache

import (
	"time"
)

// DefaultClearDuration describes the sleep time of daemon thread
var DefaultClearDuration = 15 * time.Minute

// DeleteCache del the key in memory cache
func (c *MemoryCache) DeleteCache(key interface{}) error {
	c.storage.Store(key, nil)
	c.storage.Delete(key)
	return nil
}

// Daemon del key out of date
func (c *MemoryCache) Daemon() {
	for {
		needDel := []interface{}{}
		c.storage.Range(func(key, value interface{}) bool {
			if value.(*CachedDetails).Expire.Unix() <= time.Now().Unix() {
				needDel = append(needDel, key)
			}
			return true
		})
		for _, key := range needDel {
			c.DeleteCache(key)
		}
		time.Sleep(DefaultClearDuration)
	}
}
