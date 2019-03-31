package model

type UserMatch struct {
	UserId  int64 `xorm:"not null pk BIGINT(64)"`
	MatchId int64 `xorm:"not null pk index BIGINT(64)"`
	Result  int   `xorm:"not null TINYINT(4)"`
	Rank    int   `xorm:"not null SMALLINT(4)"`
}
