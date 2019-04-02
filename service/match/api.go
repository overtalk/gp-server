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
	problems, err := m.db.GetAllProblems()
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems for intelligent compose : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all problems for intelligent compose"
		return resp
	}

	paperProblems, err := m.IntelligentCompose(problems, req.Paper)
	if err != nil {
		logger.Sugar.Errorf("failed to intelligent compose : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to intelligent compose"
		return resp
	}
	paper.P = paperProblems

	if err := m.db.AddMatch(paper, match); err != nil {
		logger.Sugar.Errorf("failed to new paper : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to new paper"
		return resp
	}
	resp.Result = true
	return resp
}

// GetMatches : get all matches
func (m *Match) GetMatches(r *http.Request) proto.Message {
	return nil
}

// GetMatchByID : get match by id
func (m *Match) GetMatchByID(r *http.Request) proto.Message {
	return nil
}

// GetMatchPaper : get the paper info
func (m *Match) GetMatchPaper(r *http.Request) proto.Message {
	return nil
}

// EditMatch : edit match
func (m *Match) EditMatch(r *http.Request) proto.Message {
	return nil
}
