package model

type TestData struct {
	Id        int64  `xorm:"pk autoincr BIGINT(64)"`
	ProblemId int64  `xorm:"not null index BIGINT(64)"`
	In        string `xorm:"TEXT"`
	Out       string `xorm:"TEXT"`
}
