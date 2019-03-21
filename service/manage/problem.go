package manage

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/gorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/file"
)

// GetProblems : get problems
func (m *BackStageManage) GetProblems(r *http.Request) proto.Message {
	req := &protocol.GetProblemsReq{}
	resp := &protocol.GetProblemsResp{Status: &protocol.Status{}}
	// get token and data
	data, token, err := getReqAndToken(r)
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
	_, err = m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	problems, err := m.db.GetProblems(req.Tag)
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "ailed to get problems list"
		return resp
	}

	for _, problem := range problems {
		resp.Problems = append(resp.Problems, problem.TurnMinProto())
	}
	return resp
}

// GetProblemByID : get problem by id
func (m *BackStageManage) GetProblemByID(r *http.Request) proto.Message {
	req := &protocol.GetProblemByIDReq{}
	resp := &protocol.GetProblemByIDResp{Status: &protocol.Status{}}
	// get token and data
	data, token, err := getReqAndToken(r)
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
	_, err = m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	problem, err := m.db.GetProblemByID(req.Id)
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
func (m *BackStageManage) AddProblem(r *http.Request) proto.Message {
	req := &protocol.AddProblemReq{}
	resp := &protocol.AddProblemResp{Status: &protocol.Status{}}

	status := m.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// add problem
	p := model.TurnProblem(req.Problem)
	relativePath := m.judgeFilePath + getJudgeFileRelativePath(p.Title)
	if err := file.Write(relativePath+"_in.txt", req.Problem.JudgeInFile); err != nil {
		resp.IsSuccess = false
		return resp
	}
	if err := file.Write(relativePath+"_out.txt", req.Problem.JudgeOutFile); err != nil {
		resp.IsSuccess = false
		return resp
	}

	p.JudgeFile = relativePath
	if err := m.db.AddProblem(p); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}

// EditProblem : edit problem to db
func (m *BackStageManage) EditProblem(r *http.Request) proto.Message {
	req := &protocol.EditProblemReq{}
	resp := &protocol.EditProblemResp{Status: &protocol.Status{}}

	status := m.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	if err := m.db.UpdateProblem(model.TurnProblem(req.Problem)); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}
