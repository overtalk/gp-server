package auth

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

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

	token, err := a.cache.UpdateToken(user.ID)
	if err != nil {
		logger.Sugar.Errorf("failed to update token for player[%d]", user.ID)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to generate token"
		return resp
	}

	resp.User = user.TurnProto()
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
