package db

import (
	"github.com/qinhan-shu/gp-server/model/transform"
)

// GetGlobalAnnouncements : get global announcements
func (m *MysqlDriver) GetGlobalAnnouncements() ([]*transform.AnnouncementWithName, error) {
	announcements := make([]*transform.AnnouncementWithName, 0)
	if err := m.conn.Join("INNER", "user", "user.id = announcement.publisher").
		Where("class_id is null").
		Find(&announcements); err != nil {
		return nil, err
	}
	return announcements, nil
}

// GetAnnouncementsByClassID : get all announcements of a certain class
func (m *MysqlDriver) GetAnnouncementsByClassID(classID int64) ([]*transform.AnnouncementWithName, error) {
	announcements := make([]*transform.AnnouncementWithName, 0)
	if err := m.conn.Join("INNER", "user", "user.id = announcement.publisher").
		Where("class_id = ?", classID).
		Find(&announcements); err != nil {
		return nil, err
	}
	return announcements, nil
}
