package manage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/gorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/file"
)

// GetProblems : get problems
func (m *BackStageManage) GetProblems(c *gin.Context) (int, interface{}) {
	// get request and response
	code := http.StatusOK
	req := &protocol.GetProblemsReq{}
	resp := &protocol.GetProblemsResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	_, err = m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	problems, err := m.db.GetProblems(req.Tag)
	if err != nil {
		logger.Sugar.Errorf("failed to get all problems : %v", err)
		code = http.StatusInternalServerError
		return code, resp
	}

	for _, problem := range problems {
		resp.Problems = append(resp.Problems, problem.TurnMinProto())
	}
	return code, resp
}

// GetProblemByID : get problem by id
func (m *BackStageManage) GetProblemByID(c *gin.Context) (int, interface{}) {
	// get request and response
	code := http.StatusOK
	req := &protocol.GetProblemByIDReq{}
	resp := &protocol.GetProblemByIDResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	_, err = m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	problem, err := m.db.GetProblemByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get problem by id : %v", err)
		code = http.StatusInternalServerError
		return code, resp
	}

	resp.Problem = problem.TurnProto()
	return code, resp
}

// AddProblem : add problem to db
func (m *BackStageManage) AddProblem(c *gin.Context) (int, interface{}) {
	// get request and response
	code := http.StatusOK
	req := &protocol.AddProblemReq{}
	resp := &protocol.AddProblemResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusUnauthorized
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to add problem[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	// add problem
	p := model.TurnProblem(req.Problem)
	relativePath := m.judgeFilePath + getJudgeFileRelativePath(p.Title)
	if err := file.Write(relativePath+"_in.txt", req.Problem.JudgeInFile); err != nil {
		resp.IsSuccess = false
		return code, resp
	}
	if err := file.Write(relativePath+"_out.txt", req.Problem.JudgeOutFile); err != nil {
		resp.IsSuccess = false
		return code, resp
	}

	p.JudgeFile = relativePath
	if err := m.db.AddProblem(p); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return code, resp
}

// EditProblem : edit problem to db
func (m *BackStageManage) EditProblem(c *gin.Context) (int, interface{}) {
	// get request and response
	code := http.StatusOK
	req := &protocol.EditProblemReq{}
	resp := &protocol.EditProblemResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to edit problem[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	if err := m.db.UpdateProblem(model.TurnProblem(req.Problem)); err != nil {
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return code, resp
}
