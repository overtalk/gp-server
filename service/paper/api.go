package paper

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
)

// NewPaper : create a new paper
func (p *Paper) NewPaper(r *http.Request) proto.Message {
	req := &protocol.NewPaperReq{}
	resp := &protocol.NewPaperResp{Status: &protocol.Status{}}

	_, status := p.checkArgsandAuth(r, req, module.TEACHER)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	problems, err := p.db.GetAllProblems()
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems for intelligent compose : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all problems for intelligent compose"
		return resp
	}

	paper := transform.ProtoToPaper(req.Paper)
	paperProblems, err := p.compose.IntelligentCompose(problems, req.Paper)
	if err != nil {
		logger.Sugar.Errorf("failed to intelligent compose : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to intelligent compose"
		return resp
	}
	paper.P = paperProblems
	for _, v := range  paperProblems {
		problem, err := p.db.GetProblemByID(v.ProblemId)
		if err != nil {
			logger.Sugar.Errorf("failed to get problem by id after compose : %v", err)
			resp.Status.Code = protocol.Code_INTERNAL
			resp.Status.Message = "failed to get problem by id after compose"
			return resp
		}
		paper.ProblemsDetail = append(paper.ProblemsDetail, problem)
	}

	if err := p.db.AddPaper(paper); err != nil {
		logger.Sugar.Errorf("failed to save intelligent compose paper : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to save intelligent compose paper"
		return resp
	}

	resp.IsSuccess = true
	resp.Paper = paper.ToProto()
	return resp
}

// ModifyPaper : modify paper
func (p *Paper) ModifyPaper(r *http.Request) proto.Message {
	req := &protocol.ManualModifyPaperReq{}
	resp := &protocol.ManualModifyPaperResp{Status: &protocol.Status{}}

	_, status := p.checkArgsandAuth(r, req, module.TEACHER)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	var err error
	switch req.ModifyType {
	case protocol.ManualModifyPaperReq_ADD:
		{
			err = p.db.AddPaperProblem(req.PaperId, req.ProblemId)
		}
	default:
		{
			err = p.db.DelPaperProblem(req.PaperId, req.ProblemId)
		}
	}

	if err != nil {
		logger.Sugar.Errorf("failed to add/delete paper problem : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to add/delete paper problem"
		return resp
	}

	resp.IsSuccess = true
	return resp
}
