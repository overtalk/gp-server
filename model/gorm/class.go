package model

// Class : table `class`
type Class struct {
	ID          int    `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Tutor       int    `gorm:"type : int(20); not null"`
	Name        string `gorm:"type :  varchar(100); not null"`
	CreatedTime int64  `gorm:"type : int(64); not null"`
}
