package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetGlobalAnnouncementsNum(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	num, err := mysqlDriver.GetGlobalAnnouncementsNum()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("the num of global announcements : %d", num)
}

func TestMysqlDriver_GetGlobalAnnouncements(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var pageNum int64 = 3
	var pageIndex int64 = 1
	announcements, err := mysqlDriver.GetGlobalAnnouncements(pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}
	for _, announcement := range announcements {
		t.Log(announcement)
	}
}

func TestMysqlDriver_GetAnnouncementsByClassID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var classID int64 = 1
	announcements, err := mysqlDriver.GetAnnouncementsByClassID(classID)
	if err != nil {
		t.Error(err)
		return
	}
	for _, announcement := range announcements {
		t.Log(announcement)
	}
}

func TestMysqlDriver_GetAnnouncementDetail(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 21
	announcement, err := mysqlDriver.GetAnnouncementDetail(id)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", announcement)
}

func TestMysqlDriver_AddGlobalAnnouncement(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	announcement := &model.Announcement{
		Title:          "sdfs",
		Detail:         "sdfsfsdfsfd",
		Publisher:      1,
		LastUpdateTime: time.Now().Unix(),
		CreateTime:     time.Now().Unix(),
	}
	if err := mysqlDriver.AddAnnouncement(announcement); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", announcement)
}

func TestMysqlDriver_EditGlobalAnnouncement(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 24
	origin, err := mysqlDriver.GetAnnouncementDetail(id)
	if err != nil {
		t.Error(err)
		return
	}

	change := &model.Announcement{
		Id:     id,
		Detail: origin.Detail + "000",
	}
	if err := mysqlDriver.EditAnnouncement(change); err != nil {
		t.Error(err)
		return
	}

	changed, err := mysqlDriver.GetAnnouncementDetail(id)
	if err != nil {
		t.Error(err)
		return
	}

	if !assert.NotEqual(t, origin.Detail, changed.Detail) {
		t.Error("filed [Detail] is not changed")
		return
	}

	if !assert.Equal(t, changed.Detail, change.Detail) {
		t.Error("filed [Detail] is not changed to expected value")
		return
	}

	if !assert.Equal(t, origin.Title, changed.Title) {
		t.Error("other filed [Title] is changed")
		return
	}
}

func TestMysqlDriver_DelAnnouncement(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	newAnnouncement := &model.Announcement{
		Title:          "sdfsasdfasdf",
		Detail:         "sdfffffffffff",
		Publisher:      1,
		LastUpdateTime: time.Now().Unix(),
		CreateTime:     time.Now().Unix(),
	}
	if err := mysqlDriver.AddAnnouncement(newAnnouncement); err != nil {
		t.Error(err)
		return
	}

	if err := mysqlDriver.DelAnnouncement(newAnnouncement.Id); err != nil {
		t.Error(err)
		return
	}

	_, err = mysqlDriver.GetAnnouncementDetail(newAnnouncement.Id)
	if err == nil {
		t.Error("failed to delete announcement")
		return
	}
}

func TestAddSomeGlobalAnnouncement(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 10; i++ {
		announcement := &model.Announcement{
			Title:          "测试全局公告" + fmt.Sprintf("%d", i),
			Detail:         "测试全局公告" + fmt.Sprintf("%d的内容", i),
			Publisher:      1,
			LastUpdateTime: time.Now().Unix(),
			CreateTime:     time.Now().Unix(),
		}
		if err := mysqlDriver.AddAnnouncement(announcement); err != nil {
			t.Error(err)
			return
		}

		t.Logf("%+v\n", announcement)
	}

}
