package cache

import (
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
	return &MemoryCache{}
}
