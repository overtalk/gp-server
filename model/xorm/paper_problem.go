package model

type PaperProblem struct {
	ProblemId int64 `xorm:"not null pk BIGINT(64)"`
	PaperId   int64 `xorm:"not null pk index BIGINT(64)"`
	Index     int   `xorm:"not null pk INT(11)"`
}
