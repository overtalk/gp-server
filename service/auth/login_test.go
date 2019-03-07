package auth_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/auth"
	"github.com/qinhan-shu/gp-server/service/config"
)

func TestAuth_LoginAndLogOut(t *testing.T) {
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	authModule := auth.NewAuth(dataStorage)

	// login operations
	loginReqBytes, err := proto.Marshal(&protocol.LoginReq{
		Username: "aaa",
		Password: "aaa",
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Request: loginReqBytes,
	}
	data := authModule.Login(args)
	loginResp := data.(*protocol.LoginResp)
	if loginResp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}
	t.Log(loginResp)

	// check the redis
	userID, err := dataStorage.Cache.GetUserIDByToken(loginResp.Token)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("userID = %d", userID)

	// logout operations
	logoutReqBytes, err := proto.Marshal(&protocol.LogOutReq{
		Username: "aaa",
	})
	if err != nil {
		t.Error(err)
		return
	}
	args = map[string]interface{}{
		module.Request: logoutReqBytes,
		module.Token:   loginResp.Token,
	}
	data = authModule.Logout(args)
	logoutResp := data.(*protocol.LogOutResp)
	if logoutResp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}
	t.Log("logoutResp.Code = ", logoutResp.Code)

	// check the redis again
	userID, err = dataStorage.Cache.GetUserIDByToken(loginResp.Token)
	if err == nil {
		t.Errorf("logout fail. get userID[%d] using old token[%s]", userID, loginResp.Token)
		return
	}
}
