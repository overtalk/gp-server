package transform

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// Paper : paper and it's problems
type Paper struct {
	model.Paper `xorm:"extends"`
	P           []*model.PaperProblem
}

// ProtoToPaper : turn protobuf to Paper
func ProtoToPaper(p *protocol.Paper) *Paper {
	paper := model.Paper{
		Difficulty: int(p.Difficulty),
		ProblemNum: int(p.ProblemNum),
	}
	return &Paper{
		Paper: paper,
	}
}
