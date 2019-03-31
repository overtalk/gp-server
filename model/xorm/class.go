package model

type Class struct {
	Id           int64  `xorm:"pk autoincr BIGINT(64)"`
	Tutor        int64  `xorm:"not null BIGINT(64)"`
	Name         string `xorm:"not null TEXT"`
	Introduction string `xorm:"TEXT"`
	Number       int    `xorm:"not null default 0 INT(11)"`
	IsCheck      int    `xorm:"not null default 0 TINYINT(1)"`
	CreateTime   int64  `xorm:"not null BIGINT(64)"`
}
