package model

// Match : table `match`
type Match struct {
	ID             int64  `gorm:"type : int(20); primary_key; not null; auto_increment"`
	PaperID        int    `gorm:"type : int(20); not null"`
	InvitationCode string `gorm:"type : varchar(50); not null; unique"`
	IsPublic       bool   `gorm:"type : boolean; not null; default : 1"`
	StartTime      int64  `gorm:"type : int(64); not null"`
	Duration       int64  `gorm:"type : int(20); not null"`
}
