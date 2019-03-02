package module

// Cache : distributed shared cache module
type Cache interface {
	Ping() error
	UpdateToken(playerID string) (string, error)
	DelToken(playerID string) error
}
