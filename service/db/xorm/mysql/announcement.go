package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// GetGlobalAnnouncementsNum : get the num of global announcements
func (m *MysqlDriver) GetGlobalAnnouncementsNum() (int64, error) {
	return m.conn.Where("class_id is null").Count(&model.Announcement{})
}

// GetGlobalAnnouncements : get global announcements
func (m *MysqlDriver) GetGlobalAnnouncements(pageNum, pageIndex int64) ([]*transform.AnnouncementWithName, error) {
	announcements := make([]*transform.AnnouncementWithName, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Join("INNER", "user", "user.id = announcement.publisher").
		Where("class_id is null").
		Find(&announcements); err != nil {
		return nil, err
	}
	return announcements, nil
}

// GetAnnouncementsByClassID : get all announcements of a certain class
// if pageNum == 0, return all
func (m *MysqlDriver) GetAnnouncementsByClassID(classID int64) ([]*transform.AnnouncementWithName, error) {
	announcements := make([]*transform.AnnouncementWithName, 0)
	if err := m.conn.
		Join("INNER", "user", "user.id = announcement.publisher").
		Where("class_id = ?", classID).
		Find(&announcements); err != nil {
		return nil, err
	}
	return announcements, nil
}
