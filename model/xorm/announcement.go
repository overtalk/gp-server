package model

type Announcement struct {
	CreateTime  int64  `xorm:"not null BIGINT(64)"`
	Detail      string `xorm:"not null TEXT"`
	DisableTime int64  `xorm:"not null BIGINT(64)"`
	Id          int64  `xorm:"pk autoincr BIGINT(64)"`
	Publisher   int64  `xorm:"not null BIGINT(64)"`
}
