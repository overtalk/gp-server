package model

type Paper struct {
	Id         int64  `xorm:"pk autoincr BIGINT(64)"`
	Difficulty string `xorm:"not null TEXT"`
	Tags       string `xorm:"not null TEXT"`
	Cognition  string `xorm:"not null TEXT"`
}
