package cache

import (
	"encoding/gob"
	"sync"
	"time"
)

// CachedData defines the cached data details
type CachedData struct {
	Data   []byte
	Expire time.Time
}

// MemoryCache defines a cache using momory
type MemoryCache struct {
	storage sync.Map
}

// NewDBCache returns a MemoryCache
func NewDBCache() *MemoryCache {
	gob.Register(time.Time{})
	return &MemoryCache{}
}
