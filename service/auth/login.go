package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// Login : authentication, and get token
func (a *Auth) Login(c *gin.Context) interface{} {
	// get request and response
	req := &protocol.LoginReq{}
	resp := &protocol.LoginResp{}
	if err := utils.CheckArgs(args, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	user, err := a.db.CheckPlayer(req.Account, utils.MD5(req.Password))
	if err != nil {
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	token, err := a.cache.UpdateToken(user.ID)
	if err != nil {
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	resp.User = user.TurnProto()
	resp.Token = token

	return resp
}

// Logout : log out, and del token
func (a *Auth) Logout(c *gin.Context) interface{} {
	resp := &protocol.LogOut{}
	if err := utils.CheckArgs(args, module.Token); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	a.cache.DelTokenByToken(parse.String(args[module.Token])) // nolint : err check

	return resp
}
