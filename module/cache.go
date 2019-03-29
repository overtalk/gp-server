package module

// Cache : distributed shared cache module
type Cache interface {
	Ping() error

	// token
	UpdateToken(userID int64) (string, error)
	GetUserIDByToken(token string) (int64, error)
	DelTokenByUserID(userID int64) error
	DelTokenByToken(token string) error

	// rank
	SetRank(*RankItem) error
	GetRank() ([]RankItem, error)
	DelRank(deleteNum int64) (int64, error)
}
