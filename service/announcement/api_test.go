package announcement_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/announcement"
	"github.com/qinhan-shu/gp-server/service/config"
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
