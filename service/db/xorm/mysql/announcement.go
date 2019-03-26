package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func (m *MysqlDriver) getAnnouncementByClassID(classID int64) ([]*model.Announcement, error) {
	announcements := make([]*model.Announcement, 0)
	if err := m.conn.Where("class_id = ?", classID).Find(&announcements); err != nil {
		return nil, err
	}
	return announcements, nil
}
