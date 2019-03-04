package model

// UserOj : teble `user_oj`
type UserOj struct {
	UserID      int    `gorm:"type : int(20); not null"`
	OjID        string `gorm:"type : varchar(100); not null"`
	PassTime    int    `gorm:"type : int(20); not null; default : 0"`
	RefusedTime int    `gorm:"type : int(20); not null; default : 0"`
}
