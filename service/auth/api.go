package auth

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// Register : create a new player
func (a *Auth) Register(r *http.Request) proto.Message {
	req := new(protocol.RegisterReq)
	resp := &protocol.RegisterResp{Status: &protocol.Status{}}
	data, err := utils.GetRequestBody(r)
	if err != nil {
		logger.Sugar.Infof("missing request data : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "data lose in request body"
		return resp
	}

	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	user := transform.ProtoToUser(req.User)

	if err := a.db.CreatePlayer(user); err != nil {
		logger.Sugar.Errorf("failed to create player : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to create player"
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}
	return resp
}

// Login : authentication, and get token
func (a *Auth) Login(r *http.Request) proto.Message {
	req := new(protocol.LoginReq)
	resp := &protocol.LoginResp{Status: &protocol.Status{}}
	data, err := utils.GetRequestBody(r)
	if err != nil {
		logger.Sugar.Infof("missing request data : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "data lose in request body"
		return resp
	}

	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	user, err := a.db.CheckPlayer(req.Account, utils.MD5(req.Password))
	if err != nil {
		logger.Sugar.Infof("unmatch account[%s] and pwd[%s]", req.Account, req.Password)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "unmatch account and pwd"
		return resp
	}

	token, err := a.cache.UpdateToken(user.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to update token for player[%d]", user.Id)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to generate token"
		return resp
	}

	resp.User = transform.UserToProto(user)
	resp.Token = token
	return resp
}

// Logout : log out, and del token
func (a *Auth) Logout(r *http.Request) proto.Message {
	resp := &protocol.LogoutResp{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	a.cache.DelTokenByToken(token) // nolint : err check
	return resp
}

// GetConfig : get all config
func (a *Auth) GetConfig(r *http.Request) proto.Message {
	resp := &protocol.Config{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	difficulty, err := a.db.GetAllDifficulty()
	if err != nil {
		logger.Sugar.Errorf("failed to get all difficulty : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all difficulty"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range difficulty {
		d[int64(v.Id)] = v.Detail
	}
	resp.Difficulty = d

	tags, err := a.db.GetAllTag()
	if err != nil {
		logger.Sugar.Errorf("failed to get all tags : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all tags"
		return resp
	}
	t := make(map[int64]string)
	for _, v := range tags {
		t[int64(v.Id)] = v.Detail
	}
	resp.Tags = t

	return resp
}

// GetUserRole : get user role
func (a *Auth) GetUserRole(r *http.Request) proto.Message {
	resp := &protocol.UserRole{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	roles, err := a.db.GetAllRole()
	if err != nil {
		logger.Sugar.Errorf("failed to get all difficulty : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all difficulty"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range roles {
		d[int64(v.Id)] = v.Detail
	}
	resp.Role = d

	return resp
}

// GetAllLanguage : get all languages
func (a *Auth) GetAllLanguage(r *http.Request) proto.Message {
	resp := &protocol.JudgeLanguage{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	languages, err := a.db.GetAllLanguage()
	if err != nil {
		logger.Sugar.Errorf("failed to get all languages : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all languages"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range languages {
		d[int64(v.Id)] = v.Detail
	}
	resp.Language = d

	return resp
}

// GetJudgeResult : get all judge result
func (a *Auth) GetJudgeResult(r *http.Request) proto.Message {
	resp := &protocol.JudgeResults{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	results := make(map[int64]string)

	results[0] = "SUCCESS"
	results[-1] = "WRONG_ANSWER"
	results[1] = "CPU_TIME_LIMIT_EXCEEDED"
	results[2] = "REAL_TIME_LIMIT_EXCEEDED"
	results[3] = "MEMORY_LIMIT_EXCEEDED"
	results[4] = "RUNTIME_ERROR"
	results[5] = "SYSTEM_ERROR"
	results[6] = "COMPILE_ERROR"

	resp.JudgeResults = results

	return resp
}
