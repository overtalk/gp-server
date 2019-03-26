package model

type UserClass struct {
	ClassId         int64 `xorm:"not null pk index BIGINT(64)"`
	IsAdministrator int   `xorm:"not null default 0 TINYINT(1)"`
	IsChecked       int   `xorm:"not null default 0 TINYINT(1)"`
	UserId          int64 `xorm:"not null pk BIGINT(64)"`
}
