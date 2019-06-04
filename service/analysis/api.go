package analysis

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/qinhan-shu/gp-server/logger"
	// "github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// DifficultyAnalysis : analyse data by difficulty
func (a *Analysis) DifficultyAnalysis(r *http.Request) proto.Message {
	req := &protocol.AnalysisByDifficultyReq{}
	resp := &protocol.AnalysisByDifficultyResp{Status: &protocol.Status{}}

	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		logger.Sugar.Error(err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = err.Error()
		return resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	line, pie, err := a.db.GetDifficultyAnalysis(req.UserId, req.StartTime, req.EndTime)
	if err != nil {
		logger.Sugar.Errorf("failed to get difficulty analysis : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get difficulty analysis"
		return resp
	}
	resp.Line = line
	resp.Pie = pie
	return resp
}

// TagsAnalysis : analyse data by tags
func (a *Analysis) TagsAnalysis(r *http.Request) proto.Message {
	req := &protocol.AnalysisByTagsReq{}
	resp := &protocol.AnalysisByTagsResp{Status: &protocol.Status{}}

	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		logger.Sugar.Error(err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = err.Error()
		return resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	line, pie, err := a.db.GetTagsAnalysis(req.UserId, req.StartTime, req.EndTime, req.Tags)
	if err != nil {
		logger.Sugar.Errorf("failed to get tags analysis : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get tags analysis"
		return resp
	}
	resp.Line = line
	resp.Pie = pie
	return resp
}
