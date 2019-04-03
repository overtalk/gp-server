package announcement_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/announcement"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetAnnouncements(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := announcement.NewAnnouncement(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.AnnouncementsReq{
		PageIndex: 1,
		PageNum:   3,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetAnnouncements(r)
	resp := data.(*protocol.AnnouncementsResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("total = %d", resp.Total)
	t.Logf("page index = %d", resp.PageIndex)
	t.Logf("page num = %d", resp.PageNum)
	for _, announcement := range resp.Announcements {
		t.Logf("%+v", announcement)
	}
}

func TestUserManage_GetDetail(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := announcement.NewAnnouncement(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.AnnouncementDetailReq{
		Id: 21,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetDetail(r)
	resp := data.(*protocol.AnnouncementDetailResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.Announcement)
}

func TestUserManage_AddAnnouncement(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := announcement.NewAnnouncement(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.AddAnnouncementReq{
		Announcement: &protocol.Announcement{
			Title:  "skfjlask哈哈哈",
			Detail: "asdfasd",
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.AddAnnouncement(r)
	resp := data.(*protocol.AddAnnouncementResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.IsSuccess)
}

func TestUserManage_EditAnnouncement(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := announcement.NewAnnouncement(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.EditAnnouncementReq{
		Announcement: &protocol.Announcement{
			Id:     28,
			Detail: "修改了啊",
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.EditAnnouncement(r)
	resp := data.(*protocol.EditAnnouncementResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.IsSuccess)
}

func TestUserManage_DelAnnouncement(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := announcement.NewAnnouncement(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.DelAnnouncementReq{
		Id: 28,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.DelAnnouncement(r)
	resp := data.(*protocol.DelAnnouncementResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.IsSuccess)
}
