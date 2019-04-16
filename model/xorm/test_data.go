package model

type TestData struct {
	In        string `xorm:"TEXT"`
	Index     int64  `xorm:"not null pk BIGINT(64)"`
	Out       string `xorm:"TEXT"`
	ProblemId int64  `xorm:"not null pk index BIGINT(64)"`
}
