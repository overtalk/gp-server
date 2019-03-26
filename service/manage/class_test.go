package manage_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/manage"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetClasses(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetClassesReq{
		PageIndex: 1,
		PageNum:   3,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := managerModule.GetClasses(r)
	resp := data.(*protocol.GetClassesResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	for _, class := range resp.Classes {
		t.Logf("%+v", class)
	}
}

func TestUserManage_GetClassByID(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetClassByIDReq{
		Id: 1,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := managerModule.GetClassByID(r)
	resp := data.(*protocol.GetClassByIDResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("%+v", resp.Class.Name)
	for _, announcement := range resp.Class.Announcements {
		t.Logf("%+v", announcement.Detail)
	}
}

func TestUserManage_AddClass(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.AddClassReq{
		Class: &protocol.Class{
			Name:         "测试班级007",
			Introduction: "测试班级007的简单介绍",
			IsCheck:      false,
			Announcements: []*protocol.Announcement{
				&protocol.Announcement{
					Detail: "测试班级007的公告",
				},
			},
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := managerModule.AddClass(r)
	resp := data.(*protocol.AddClassResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("%+v", resp.IsSuccess)
}

func TestUserManage_UpdateClass(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.EditClassReq{
		Class: &protocol.Class{
			Id:   12,
			Name: "测试班级007(编辑时候被修改啦asdfasdfasdf)",
			Announcements: []*protocol.Announcement{
				&protocol.Announcement{
					Detail: "测试班级007的公告(编辑时候产生asdfasdfasd)",
				},
			},
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := managerModule.EditClass(r)
	resp := data.(*protocol.EditClassResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("%+v", resp.IsSuccess)
}
