package model

// UserClass : teble `user`
type UserClass struct {
	UserID       int64  `gorm:"type : int(20); not null"`
	ClassID      int    `gorm:"type : int(20); not null"`
	announcement []byte `gorm:"type : blob; default null"`
}
