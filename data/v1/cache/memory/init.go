package cache

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
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

