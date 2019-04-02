package db_test

import (
	"testing"
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
