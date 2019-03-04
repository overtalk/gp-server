package model

// UserMatch : teble `user`
type UserMatch struct {
	UserID  int `gorm:"type : int(20); not null"`
	MatchID int `gorm:"type : int(20); not null"`
	result  int `gorm:"type : tinyint(4); not null"`
	rank    int `gorm:"type : smallint(4); not null"`
}
