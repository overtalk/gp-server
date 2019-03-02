package model

import (
	"time"
)

// Test : teble `test`
type Test struct {
	ID          int       `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Nickname    string    `gorm:"type : varchar(50); not null; default : 'nickname'"`
	CreatedAt   time.Time `gorm:"type : timestamp; not null"`
	Achievement []byte    `gorm:"type : blob"`
	Level       int       `gorm:"type : int(11); not null; default : 1"`
}
