package auth

import (
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// Login : authentication, and get token
func (a *Auth) Login(data ...interface{}) interface{} {
	req := &protocol.LoginReq{}
	resp := &protocol.LoginResp{}
	if len(data) != 1 {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	if err := proto.Unmarshal(parse.Bytes(data[0]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// TODO: check the player, and get playerID
	playerID := "test"

	token, err := a.cache.UpdateToken(playerID)
	if err != nil {
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	resp.Token = token
	return resp
}

// Logout : log out, and del token
func (a *Auth) Logout(data ...interface{}) interface{} {
	return nil
}
