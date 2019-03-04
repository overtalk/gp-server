package model

// User : teble `user`
type User struct {
	ID            int    `gorm:"type : int(20); primary_key; not null; auto_increment"`
	AuthCode      string `gorm:"type : varchar(50); not null; unique"`
	OperationAuth int    `gorm:"type : tinyint(4); not null; default : 0"`
	Role          int    `gorm:"type : tinyint(4); not null; default : 0"`
	Name          string `gorm:"type : varchar(50); not null"`
	Sex           bool   `gorm:"type : boolean; not null"`
	Email         string `gorm:"type : varchar(50); not null"`
	Academy       string `gorm:"type : varchar(50); not null"`
	Major         string `gorm:"type : varchar(50); not null"`
	LastLogin     int64  `gorm:"type : int(64); not null"`
}
