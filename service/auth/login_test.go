package auth_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/auth"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/utils"
)

func TestAuth_LoginAndLogOut(t *testing.T) {
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	authModule := auth.NewAuth(dataStorage)

	r, err := utils.MockHTTPReq("POST", "", &protocol.LoginReq{
		Account:  "jack0",
		Password: "jack0",
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := authModule.Login(r)
	resp := data.(*protocol.LoginResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
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
	r1, err := utils.MockHTTPReq("GET", resp.Token, &protocol.LoginReq{
		Account:  "jack0",
		Password: "jack0",
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}
	data = authModule.Logout(r1)
	logoutResp := data.(*protocol.LogoutResp)
	if logoutResp.Status.Code != protocol.Code_OK {
		t.Errorf("logoutResp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(logoutResp.Status.Code)])
		return
	}
	t.Log("LogOut.Code = ", logoutResp.Status.Code)

	// check the redis again
	userID, err = dataStorage.Cache.GetUserIDByToken(resp.Token)
	if err == nil {
		t.Errorf("logout fail. get userID[%d] using old token[%s]", userID, resp.Token)
		return
	}
}
