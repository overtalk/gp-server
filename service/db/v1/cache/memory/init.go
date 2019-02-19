package cache

import (
	"sync"
	"time"
)

// CachedDetails defines the cached data details
type CachedDetails struct {
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
