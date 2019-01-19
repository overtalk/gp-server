package cache

import "time"

func (c *Cache) Delete(key string) error {
	c.storage.Store(key, nil)
	c.storage.Delete(key)
	return nil
}

func (c *Cache) Daemon() {
	for {
		c.storage.Range(func(key, value interface{}) bool {
			if value.(*CacheData).Expire.Unix() <= time.Now().Unix() {
				_ = c.Delete(key.(string))
			}
			return true
		})
		time.Sleep(15 * time.Minute)
	}
}
