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
	tags, _ := json.Marshal(p.KnowledgePoints)
	cognition, _ := json.Marshal(p.Cognition)
	difficulty, _ := json.Marshal(p.Difficulty)

	paper := model.Paper{
		Tags:       string(tags),
		Cognition:  string(cognition),
		Difficulty: string(difficulty),
	}
	return &Paper{
		Paper: paper,
	}
}

// ToProto : turn paper to proto
func (p *Paper) ToProto() *protocol.Paper {
	tags := make(map[int64]int64)
	json.Unmarshal([]byte(p.Tags), &tags)

	cognition := make(map[int64]int64)
	json.Unmarshal([]byte(p.Cognition), &cognition)

	difficulty := make(map[int64]int64)
	json.Unmarshal([]byte(p.Difficulty), &difficulty)

	paper := &protocol.Paper{
		Id:              p.Id,
		Difficulty:      difficulty,
		KnowledgePoints: tags,
		Cognition:       cognition,
	}

	for _, problem := range p.ProblemsDetail {
		paper.Problems = append(paper.Problems, problem.TurnMinProto())
	}

	return paper
}
