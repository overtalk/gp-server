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

// GetAnnouncementDetail : get the detail of global Announcement
func (m *MysqlDriver) GetAnnouncementDetail(id int64) (*transform.AnnouncementWithName, error) {
	announcement := new(transform.AnnouncementWithName)
	ok, err := m.conn.
		Id(id).
		Join("INNER", "user", "user.id = announcement.publisher").
		Get(announcement)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}
	return announcement, nil
}

// AddAnnouncement : add announcement
func (m *MysqlDriver) AddAnnouncement(announcement *model.Announcement) error {
	var (
		err      error
		affected int64
	)
	if announcement.ClassId == 0 {
		affected, err = m.conn.Omit("class_id").Insert(announcement)
	} else {
		affected, err = m.conn.Insert(announcement)
	}
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// EditAnnouncement : edit announcement
func (m *MysqlDriver) EditAnnouncement(announcement *model.Announcement) error {
	affected, err := m.conn.Id(announcement.Id).Update(announcement)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// DelAnnouncement : del announcement
func (m *MysqlDriver) DelAnnouncement(id int64) error {
	affected, err := m.conn.Id(id).Delete(&model.Announcement{})
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
