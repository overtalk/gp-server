package module

// Cache : distributed shared cache module
type Cache interface {
	Ping() error
	UpdateToken(userID int) (string, error)
	DelToken(userID int) error
	GetUserIDByToken(token string) (int, error)
}
