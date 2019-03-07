package module

// Cache : distributed shared cache module
type Cache interface {
	Ping() error
	UpdateToken(userID int) (string, error)
	GetUserIDByToken(token string) (int, error)
	DelTokenByUserID(userID int) error
	DelTokenByToken(token string) error
}
