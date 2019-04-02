package model

type Paper struct {
	Difficulty     int    `xorm:"not null TINYINT(4)"`
	Id             int64  `xorm:"pk autoincr BIGINT(64)"`
	KnowledgePoint string `xorm:"not null TEXT"`
	ProblemNum     int    `xorm:"not null TINYINT(4)"`
	Tags           string `xorm:"not null TEXT"`
}
