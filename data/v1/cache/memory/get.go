package cache

import "errors"

func (c *Cache) GetAll(key string) (map[string]interface{} , error) {
	data, ok := c.storage.Load(key)
	if ok {
		return data.(*CacheData).Data, nil
	} else {
		return nil, errors.New("data not find")
	}
}
