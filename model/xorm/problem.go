package model

type Problem struct {
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	Title          string `xorm:"not null TEXT"`
	Description    string `xorm:"not null TEXT"`
	InDescription  string `xorm:"not null TEXT"`
	OutDescription string `xorm:"not null TEXT"`
	Hint           string `xorm:"TEXT"`
	JudgeLimitTime int    `xorm:"not null INT(11)"`
	JudgeLimitMem  int    `xorm:"not null INT(11)"`
	Difficulty     int    `xorm:"not null default 0 TINYINT(4)"`
	LastUsed       int64  `xorm:"not null default 0 BIGINT(64)"`
	UsedTime       int    `xorm:"not null default 0 INT(64)"`
	SubmitTime     int    `xorm:"not null default 0 INT(64)"`
	Ac             int    `xorm:"not null default 0 INT(64)"`
	Wa             int    `xorm:"not null default 0 INT(64)"`
	Tle            int    `xorm:"not null default 0 INT(64)"`
	Mle            int    `xorm:"not null default 0 INT(64)"`
	Pe             int    `xorm:"not null default 0 INT(64)"`
	Ce             int    `xorm:"not null default 0 INT(64)"`
	JudgeFile      string `xorm:"not null VARCHAR(100)"`
}
