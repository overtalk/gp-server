package model

type Match struct {
	Duration     int    `xorm:"not null INT(20)"`
	Id           int64  `xorm:"pk autoincr BIGINT(64)"`
	Introduction string `xorm:"TEXT"`
	IsPublic     int    `xorm:"not null default 1 TINYINT(1)"`
	PaperId      int64  `xorm:"not null BIGINT(64)"`
	StartTime    int64  `xorm:"not null BIGINT(64)"`
	Title        string `xorm:"not null TEXT"`
}
