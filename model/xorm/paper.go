package model

type Paper struct {
	Difficulty     int    `xorm:"not null TINYINT(4)"`
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	KnowledgePoint string `xorm:"not null TEXT"`
	Problems       string `xorm:"not null JSON"`
}
