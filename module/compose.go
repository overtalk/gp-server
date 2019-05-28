package module

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// Compose : paper compose module
type Compose interface {
	IntelligentCompose(problems []*model.Problem, paper *protocol.Paper) ([]*model.PaperProblem, error)
}
