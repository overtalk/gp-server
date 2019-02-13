package model

// Cache defines the service cache model (redis ...)
// all distributed node use one
type Cache interface {
	Ping() error
	UpdateToken(playerID string) (string, error)
	DelToken(playerID string) error
}
