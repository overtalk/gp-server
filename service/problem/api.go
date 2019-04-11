package problem

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetProblems : get problems
func (p *Problem) GetProblems(r *http.Request) proto.Message {
	req := &protocol.GetProblemsReq{}
	resp := &protocol.GetProblemsResp{Status: &protocol.Status{}}
	// get token and data
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
	_, err = p.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	var problems []*transform.IntactProblem
	if req.GetAll {
		problems, err = p.db.GetProblems(req.PageNum, req.PageIndex)
	} else {
		problems, err = p.db.GetProblemsByTagID(req.PageNum, req.PageIndex, int(req.Tag))
	}
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get problems list"
		return resp
	}

	for _, problem := range problems {
		resp.Problems = append(resp.Problems, problem.TurnMinProto())
	}

	// get all number
	problemsNum, err := p.db.GetProblemsNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of problems : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of problems"
		return resp
	}
	resp.Total = problemsNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum

	return resp
}

// GetProblemByID : get problem by id
func (p *Problem) GetProblemByID(r *http.Request) proto.Message {
	req := &protocol.GetProblemByIDReq{}
	resp := &protocol.GetProblemByIDResp{Status: &protocol.Status{}}
	// get token and data
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
	_, err = p.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	problem, err := p.db.GetProblemByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get problem by id : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to get problem[id = %d]", req.Id)
		return resp
	}

	resp.Problem = problem.TurnProto()
	return resp
}

// AddProblem : add problem to db
func (p *Problem) AddProblem(r *http.Request) proto.Message {
	req := &protocol.AddProblemReq{}
	resp := &protocol.AddProblemResp{Status: &protocol.Status{}}

	user, status := p.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// add problem
	pro := transform.ProtoToProblem(req.Problem)
	// relativePath := m.judgeFilePath + getJudgeFileRelativePath(p.Detail.Title)
	// if err := file.Write(relativePath+"_in.txt", req.Problem.JudgeInFile); err != nil {
	// 	resp.IsSuccess = false
	// 	return resp
	// }
	// if err := file.Write(relativePath+"_out.txt", req.Problem.JudgeOutFile); err != nil {
	// 	resp.IsSuccess = false
	// 	return resp
	// }

	// p.JudgeFile = relativePath
	pro.CreateTime = time.Now().Unix()
	pro.Publisher = user.Id
	if err := p.db.AddProblem(pro); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}

// EditProblem : edit problem to db
func (p *Problem) EditProblem(r *http.Request) proto.Message {
	req := &protocol.EditProblemReq{}
	resp := &protocol.EditProblemResp{Status: &protocol.Status{}}

	_, status := p.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	if err := p.db.UpdateProblem(transform.ProtoToProblem(req.Problem)); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}
