package cache

import (
	"encoding/gob"
	"time"
)

func Connect(info CacheInfo) (*Cache, error) {
	gob.Register(time.Time{})
	return &Cache{}, nil
}
