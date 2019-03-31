package model

type UserProblem struct {
	Id               int64  `xorm:"pk autoincr BIGINT(64)"`
	UserId           int64  `xorm:"not null index BIGINT(64)"`
	ProblemId        int64  `xorm:"not null index BIGINT(64)"`
	SubmitTime       int    `xorm:"not null INT(64)"`
	Ispass           int    `xorm:"not null TINYINT(1)"`
	RunningLangurage int    `xorm:"not null TINYINT(4)"`
	RunningTime      int    `xorm:"INT(64)"`
	RunningMem       int    `xorm:"INT(64)"`
	Code             string `xorm:"not null TEXT"`
}
