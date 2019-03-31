package model

type UserClass struct {
	UserId          int64 `xorm:"not null pk BIGINT(64)"`
	ClassId         int64 `xorm:"not null pk index BIGINT(64)"`
	IsChecked       int   `xorm:"not null default 0 TINYINT(1)"`
	IsAdministrator int   `xorm:"not null default 0 TINYINT(1)"`
}
