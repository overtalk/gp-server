package match

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
)

// NewMatch : new a match
func (m *Match) NewMatch(r *http.Request) proto.Message {
	req := &protocol.NewMatchReq{}
	resp := &protocol.NewMatchResp{Status: &protocol.Status{}}

	_, status := m.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	paper := transform.ProtoToPaper(req.Paper)
	match := transform.ProtoToMatch(req.Match)

	// TODO: create paper
	m.newPaper(paper)

	if err := m.db.AddMatch(paper, match); err != nil {
		logger.Sugar.Errorf("failed to new paper : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to new paper"
		return resp
	}
	resp.Result = true
	return resp
}
