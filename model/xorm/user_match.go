package model

type UserMatch struct {
	MatchId int64 `xorm:"not null pk BIGINT(64)"`
	Rank    int   `xorm:"not null SMALLINT(4)"`
	Result  int   `xorm:"not null TINYINT(4)"`
	UserId  int64 `xorm:"not null pk BIGINT(64)"`
}
