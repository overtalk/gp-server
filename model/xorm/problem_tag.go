package model

type ProblemTag struct {
	Id        int64 `xorm:"pk autoincr BIGINT(64)"`
	ProblemId int64 `xorm:"not null index BIGINT(64)"`
	TagId     int   `xorm:"not null index INT(11)"`
}
