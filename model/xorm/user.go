package model

type User struct {
	Account   string `xorm:"not null unique VARCHAR(50)"`
	Create    int64  `xorm:"not null BIGINT(64)"`
	Email     string `xorm:"VARCHAR(50)"`
	Id        int64  `xorm:"pk autoincr BIGINT(64)"`
	LastLogin int64  `xorm:"not null BIGINT(64)"`
	Name      string `xorm:"not null VARCHAR(50)"`
	Password  string `xorm:"not null VARCHAR(100)"`
	Phone     string `xorm:"VARCHAR(20)"`
	Role      int    `xorm:"not null default 0 TINYINT(4)"`
	School    string `xorm:"VARCHAR(50)"`
	Sex       int    `xorm:"not null default 0 TINYINT(1)"`
}
