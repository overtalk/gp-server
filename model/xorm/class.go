package model

type Class struct {
	Announcement string `xorm:"not null JSON"`
	CreateTime   int64  `xorm:"not null BIGINT(64)"`
	Id           int64  `xorm:"pk autoincr BIGINT(64)"`
	Tutor        int64  `xorm:"not null BIGINT(64)"`
}
