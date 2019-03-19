package cache

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	storage sync.Map
}

type CacheInfo struct {
}

type CacheData struct {
	Data   []byte
	Expire time.Time
}

type Pool = redis.Cmdable
