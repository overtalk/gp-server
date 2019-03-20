package model

type UserClass struct {
	Announcement []byte `xorm:"BLOB"`
	ClassId      int64  `xorm:"not null pk BIGINT(64)"`
	UserId       int64  `xorm:"not null pk BIGINT(64)"`
}
