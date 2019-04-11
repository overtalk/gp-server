package model

type Match struct {
	Id           int64  `xorm:"pk autoincr BIGINT(64)"`
	PaperId      int64  `xorm:"not null BIGINT(64)"`
	IsPublic     int    `xorm:"not null default 1 TINYINT(1)"`
	Title        string `xorm:"not null TEXT"`
	Introduction string `xorm:"TEXT"`
	StartTime    int64  `xorm:"not null BIGINT(64)"`
	EndTime      int64  `xorm:"not null BIGINT(64)"`
}
