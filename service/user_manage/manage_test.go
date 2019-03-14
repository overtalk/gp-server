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

func TestUserManage_AddUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// add users operations
	addUsersReqBytes, err := proto.Marshal(&protocol.AddUsersReq{
		Users: []*protocol.UserInfo{
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_10",
				Password: "AddUsers_10",
			},
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_20",
				Password: "AddUsers_20",
			},
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_30",
				Password: "AddUsers_30",
			},
			&protocol.UserInfo{
				Id:       2,
				Name:     "asdf",
				Sex:      false,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_40",
				Password: "AddUsers_40",
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: addUsersReqBytes,
	}

	data := userManagerModule.AddUsers(args)
	addUsersResp := data.(*protocol.AddUsersResp)
	if addUsersResp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}

	t.Log("-------Succeed-------")
	for _, user := range addUsersResp.Succeed {
		t.Logf("%+v", user)
	}

	t.Log("-------failed-------")
	for _, user := range addUsersResp.Fail {
		t.Logf("%+v", user)
	}
}
