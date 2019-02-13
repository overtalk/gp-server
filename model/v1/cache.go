package model

// Cache defines the service cache model (redis ...)
type Cache interface {
	Ping() error
	UpdateToken(playerID string) (string, error)
}
