package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	// "github.com/qinhan-shu/gp-server/model/xorm"
	// "github.com/qinhan-shu/gp-server/protocol"
)

// Paper : paper module
type Paper interface {
	// IntelligentCompose(problems []*model.Problem, paper *protocol.Paper) ([]*model.PaperProblem, error)
	NewPaper(r *http.Request) proto.Message
	ModifyPaper(r *http.Request) proto.Message
}
