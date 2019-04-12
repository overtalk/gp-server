package model

type Paper struct {
	Id         int64  `xorm:"pk autoincr BIGINT(64)"`
	Difficulty int    `xorm:"not null INT(11)"`
	ProblemNum int    `xorm:"not null INT(11)"`
	Tags       string `xorm:"not null TEXT"`
}
