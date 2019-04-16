package model

type Problem struct {
	Ac             int    `xorm:"not null default 0 INT(64)"`
	Ce             int    `xorm:"not null default 0 INT(64)"`
	CreateTime     int64  `xorm:"not null default 0 BIGINT(64)"`
	Description    string `xorm:"not null TEXT"`
	Difficulty     int    `xorm:"not null default 1 index INT(11)"`
	Hint           string `xorm:"TEXT"`
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	InDescription  string `xorm:"not null TEXT"`
	JudgeFile      string `xorm:"not null VARCHAR(100)"`
	JudgeLimitMem  int    `xorm:"not null INT(11)"`
	JudgeLimitTime int    `xorm:"not null INT(11)"`
	LastUsed       int64  `xorm:"not null default 0 BIGINT(64)"`
	Mle            int    `xorm:"not null default 0 INT(64)"`
	OutDescription string `xorm:"not null TEXT"`
	Pe             int    `xorm:"not null default 0 INT(64)"`
	Publisher      int64  `xorm:"not null index BIGINT(64)"`
	SubmitTime     int    `xorm:"not null default 0 INT(64)"`
	Tags           string `xorm:"not null TEXT"`
	Title          string `xorm:"not null TEXT"`
	Tle            int    `xorm:"not null default 0 INT(64)"`
	UsedTime       int    `xorm:"not null default 0 INT(64)"`
	Wa             int    `xorm:"not null default 0 INT(64)"`
}
