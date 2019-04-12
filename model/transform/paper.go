package transform

import (
	"encoding/json"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// Paper : paper and it's problems
type Paper struct {
	model.Paper    `xorm:"extends"`
	P              []*model.PaperProblem
	ProblemsDetail []*IntactProblem
}

// ProtoToPaper : turn protobuf to Paper
func ProtoToPaper(p *protocol.Paper) *Paper {
	tags, _ := json.Marshal(p.Tags)

	paper := model.Paper{
		Tags:       string(tags),
		Difficulty: int(p.Difficulty),
		ProblemNum: int(p.ProblemNum),
	}
	return &Paper{
		Paper: paper,
	}
}

// ToProto : turn paper to proto
func (p *Paper) ToProto() *protocol.Paper {
	tags := make([]int64, 0)
	json.Unmarshal([]byte(p.Tags), &tags)

	paper := &protocol.Paper{
		Id:         p.Id,
		Difficulty: int64(p.Difficulty),
		Tags:       tags,
		ProblemNum: int64(p.ProblemNum),
	}

	for _, problem := range p.ProblemsDetail {
		paper.Problems = append(paper.Problems, problem.TurnMinProto())
	}

	return paper
}
