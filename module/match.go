package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// Match : match module
type Match interface {
	IntelligentCompose(problems []*model.Problem, paper *protocol.Paper) ([]*model.PaperProblem, error)
	NewMatch(r *http.Request) proto.Message
	GetMatches(r *http.Request) proto.Message
	GetMatchByID(r *http.Request) proto.Message
	GetMatchPaper(r *http.Request) proto.Message
	EditMatch(r *http.Request) proto.Message
}
