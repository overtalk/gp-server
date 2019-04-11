package model

type Announcement struct {
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	Publisher      int64  `xorm:"not null BIGINT(64)"`
	Title          string `xorm:"not null TEXT"`
	Detail         string `xorm:"not null TEXT"`
	ClassId        int64  `xorm:"index BIGINT(64)"`
	CreateTime     int64  `xorm:"not null BIGINT(64)"`
	LastUpdateTime int64  `xorm:"not null BIGINT(64)"`
}
