package model

type Tag struct {
	Detail string `xorm:"not null VARCHAR(100)"`
	Id     int    `xorm:"not null pk autoincr INT(11)"`
}
