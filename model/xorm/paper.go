package model

type Paper struct {
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	Difficulty     int    `xorm:"not null TINYINT(4)"`
	KnowledgePoint string `xorm:"not null TEXT"`
	Problems       string `xorm:"not null TEXT"`
}
