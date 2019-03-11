package model

// Problem : teble `problem`
type Problem struct {
	ID          int    `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Title       string `gorm:"type : varchar(300); not null"`
	Description string `gorm:"type : text; not null"`
	Example     string `gorm:"type : text; not null"`
	JudgeFile   string `gorm:"type : varchar(100); not null"`
	JudgeLimit  string `gorm:"type : json"`
	SubmitTime  int    `gorm:"type : int(20); not null; default : 0"`
	AccpetTime  int    `gorm:"type : int(20); not null; default : 0"`
	Tags        string `gorm:"type : varchar(300); not null"`
	Difficulty  int    `gorm:"type : tinyint(4); not null; default : 0"`
	LastUsed    int64  `gorm:"type : int(64); not null"`
	UsedTime    int    `gorm:"type : int(20); not null; default : 0"`
}
