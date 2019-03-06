package module

// Cache : distributed shared cache module
type Cache interface {
	Ping() error
	UpdateToken(userID string) (string, error)
	DelToken(userID string) error
	GetUserIDByToken(token string) (string, error)
}
