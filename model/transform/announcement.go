package transform

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// AnnouncementWithName : announcement and publisher name
type AnnouncementWithName struct {
	model.Announcement `xorm:"extends"`
	Name               string
}

func (AnnouncementWithName) TableName() string {
	return "announcement"
}

// AnnouncementToProto : turn announcement to protobuf
func AnnouncementToProto(u *AnnouncementWithName) *protocol.Announcement {
	return &protocol.Announcement{
		Publisher:      u.Name,
		Title:          u.Titile,
		Detail:         u.Detail,
		CreateTime:     u.CreateTime,
		LastUpdateTime: u.LastUpdateTime,
	}
}

// ProtoToAnnouncement : turn protobuf to announcement
func ProtoToAnnouncement(a *protocol.Announcement) *model.Announcement {
	return &model.Announcement{
		Titile: a.Title,
		Detail: a.Detail,
	}
}
