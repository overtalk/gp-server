package model

// Cache defines the service cache model (redis ...)
type Cache interface {
	UpdateToken(playerID string) error
}
