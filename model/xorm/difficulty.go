package model

type Difficulty struct {
	Detail string `xorm:"not null VARCHAR(100)"`
	Id     int    `xorm:"not null pk autoincr INT(11)"`
}
