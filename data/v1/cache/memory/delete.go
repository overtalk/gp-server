package cache

import "time"

func (c *Cache) Delete(key interface{}) error {
	c.storage.Store(key, nil)
	c.storage.Delete(key)
	return nil
}

func (c *Cache) Daemon() {
	for {
		needDel := []interface{}{}
		c.storage.Range(func(key, value interface{}) bool {
			if value.(*CacheData).Expire.Unix() <= time.Now().Unix() {
				needDel = append(needDel, key)
			}
			return true
		})
		for key := range needDel {
			_ = c.Delete(key)
		}
		time.Sleep(15 * time.Minute)
	}
}
