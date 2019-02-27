package module

// Cache : 分布式共享缓存模块接口
type Cache interface {
	Ping() error
	UpdateToken(playerID string) (string, error)
	DelToken(playerID string) error
}
