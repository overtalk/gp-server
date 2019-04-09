package model

type Paper struct {
	Cognition  string `xorm:"not null TEXT"`
	Difficulty string `xorm:"not null TEXT"`
	Id         int64  `xorm:"pk autoincr BIGINT(64)"`
	Tags       string `xorm:"not null TEXT"`
}
