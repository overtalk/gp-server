package cache

// RedisConfig : config to new redis
type RedisConfig struct {
	Addr     string
	Password string
	PoolSize int
}
