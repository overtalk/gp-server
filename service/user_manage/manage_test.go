package user_manage_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/user_manage"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// get user operations
	getUsersReqBytes, err := proto.Marshal(&protocol.GetUsersReq{
		GetAll: false,
		Role:   protocol.Role_TEACHER,
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: getUsersReqBytes,
	}

	data := userManagerModule.GetUsers(args)
	getUsersResp := data.(*protocol.GetUsersResp)
	if getUsersResp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}
	for _, user := range getUsersResp.Users {
		t.Logf("%+v", user)
	}
}
