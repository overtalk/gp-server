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
		Account:  "jack0",
		Password: "jack0",
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Request: loginReqBytes,
	}
	data := authModule.Login(args)
	resp := data.(*protocol.LoginResp)
	if resp.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Code)])
		return
	}
	t.Log(resp)

	// check the redis
	userID, err := dataStorage.Cache.GetUserIDByToken(resp.Token)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("userID = %d", userID)

	// logout operations
	args = map[string]interface{}{
		module.Token: resp.Token,
	}
	data = authModule.Logout(args)
	logoutResp := data.(*protocol.LogOut)
	if logoutResp.Code != protocol.Code_OK {
		t.Errorf("logoutResp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(logoutResp.Code)])
		return
	}
	t.Log("LogOut.Code = ", logoutResp.Code)

	// check the redis again
	userID, err = dataStorage.Cache.GetUserIDByToken(resp.Token)
	if err == nil {
		t.Errorf("logout fail. get userID[%d] using old token[%s]", userID, resp.Token)
		return
	}
}
