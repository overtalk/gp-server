package match

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// IntelligentCompose : intelligent compose
func (m *Match) IntelligentCompose(problems []*model.Problem, paper *protocol.Paper) ([]*model.PaperProblem, error) {
	p := make([]*model.PaperProblem, 0)
	p = append(p, &model.PaperProblem{
		Index:     1,
		ProblemId: 1,
	})
	p = append(p, &model.PaperProblem{
		Index:     2,
		ProblemId: 1,
	})
	return p, nil
}
