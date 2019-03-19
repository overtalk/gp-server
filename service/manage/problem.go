package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/gorm"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/file"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetProblems : get problems
func (m *BackStageManage) GetProblems(c *gin.Context) interface{} {
	// get request and response
	req := &protocol.GetProblemsReq{}
	resp := &protocol.GetProblemsResp{}
	if err := utils.CheckArgs(args, module.Request, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// check token
	_, err := m.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		resp.Code = protocol.Code_INVAILD_TOKEN
		return resp
	}

	problems, err := m.db.GetProblems(req.Tag)
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems : %v", err)
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	for _, problem := range problems {
		resp.Problems = append(resp.Problems, problem.TurnMinProto())
	}
	return resp
}

// GetProblemByID : get problem by id
func (m *BackStageManage) GetProblemByID(c *gin.Context) interface{} {
	// get request and response
	req := &protocol.GetProblemByIDReq{}
	resp := &protocol.GetProblemByIDResp{}
	if err := utils.CheckArgs(args, module.Request, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// check token
	_, err := m.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		resp.Code = protocol.Code_INVAILD_TOKEN
		return resp
	}

	problem, err := m.db.GetProblemByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get problem by id : %v", err)
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	resp.Problem = problem.TurnProto()
	return resp
}

// AddProblem : add problem to db
func (m *BackStageManage) AddProblem(c *gin.Context) interface{} {
	// get request and response
	req := &protocol.AddProblemReq{}
	resp := &protocol.AddProblemResp{}
	if err := utils.CheckArgs(args, module.Request, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		resp.Code = protocol.Code_INVAILD_TOKEN
		return resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to add problem[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		resp.Code = protocol.Code_PERMISSION_DENIED
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
func (m *BackStageManage) EditProblem(c *gin.Context) interface{} {
	// get request and response
	req := &protocol.EditProblemReq{}
	resp := &protocol.EditProblemResp{}
	if err := utils.CheckArgs(args, module.Request, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		resp.Code = protocol.Code_INVAILD_TOKEN
		return resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to edit problem[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	if err := m.db.UpdateProblem(model.TurnProblem(req.Problem)); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}
