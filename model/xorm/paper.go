package model

type Paper struct {
	Difficulty int    `xorm:"not null INT(11)"`
	Id         int64  `xorm:"pk autoincr BIGINT(64)"`
	ProblemNum int    `xorm:"not null INT(11)"`
	Tags       string `xorm:"not null TEXT"`
}
