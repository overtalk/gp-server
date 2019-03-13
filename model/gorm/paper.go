package model

// Paper : table `paper`
type Paper struct {
	ID             int64  `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Difficulty     int    `gorm:"type : tinyint(4); not null"`
	KnowledgePoint string `gorm:"type : text; not null"`
}
