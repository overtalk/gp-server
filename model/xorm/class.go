package model

type Class struct {
	CreateTime   int64  `xorm:"not null BIGINT(64)"`
	Id           int64  `xorm:"pk autoincr BIGINT(64)"`
	Introduction string `xorm:"TEXT"`
	IsCheck      int    `xorm:"not null default 0 TINYINT(1)"`
	Name         string `xorm:"not null TEXT"`
	Number       int    `xorm:"not null default 0 INT(11)"`
	Tutor        int64  `xorm:"not null BIGINT(64)"`
}
