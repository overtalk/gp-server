package cache

import (
	"errors"
	"strings"

	"github.com/go-redis/redis"

	"github.com/qinhan-shu/gp-server/logger"
)

var (
	// ErrInvalidRedisAddr describes error of invalid redis address
	ErrInvalidRedisAddr = errors.New("invalid redis address")
)

// RedisCache describes redis
type RedisCache struct {
	client redis.Cmdable
}

// NewRedisCache creates a new RedisCache
func NewRedisCache(conf *RedisConfig) (*RedisCache, error) {
	if conf.Addr == "" {
		return nil, ErrInvalidRedisAddr
	}

	redisCache := new(RedisCache)
	redisAddrs := strings.Split(conf.Addr, ",")
	if len(redisAddrs) > 1 {
		redisCache.client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisAddrs,
			Password: conf.Password,
			PoolSize: conf.PoolSize,
		})
	} else {
		redisCache.client = redis.NewClient(&redis.Options{
			Addr:     conf.Addr,
			Password: conf.Password,
			PoolSize: conf.PoolSize,
		})
	}

	if _, err := redisCache.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return nil, err
	}
	return redisCache, nil
}

// Ping is to test the redis connection
func (r *RedisCache) Ping() error {
	if _, err := r.client.Ping().Result(); err != nil {
		return err
	}

	return nil
}
