package db_test

import (
	"testing"
)

func TestMysqlDriver_GetGlobalAnnouncements(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	announcements, err := mysqlDriver.GetGlobalAnnouncements()
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
