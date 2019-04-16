package model

type PaperProblem struct {
	Index     int   `xorm:"not null pk INT(11)"`
	PaperId   int64 `xorm:"not null pk index BIGINT(64)"`
	ProblemId int64 `xorm:"not null pk BIGINT(64)"`
}
