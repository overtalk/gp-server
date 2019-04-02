package transform

import (
	"time"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// IntactClass : intact class
type IntactClass struct {
	Class         *model.Class
	Announcements []*AnnouncementWithName
	TutorName     string
}

// TurnProto : turn intactProblem to protobuf
func (c *IntactClass) TurnProto() *protocol.Class {
	protobufClass := &protocol.Class{
		Tutor:        c.TutorName,
		Id:           c.Class.Id,
		Name:         c.Class.Name,
		Introduction: c.Class.Introduction,
		Number:       int64(c.Class.Number),
		IsCheck:      c.Class.IsCheck == 1,
		CreateTime:   c.Class.CreateTime,
	}
	for _, announcement := range c.Announcements {
		protobufClass.Announcements = append(protobufClass.Announcements, &protocol.Announcement{
			Publisher:      c.TutorName,
			Title:          announcement.Title,
			Detail:         announcement.Detail,
			CreateTime:     announcement.CreateTime,
			LastUpdateTime: announcement.LastUpdateTime,
		})
	}
	return protobufClass
}

// TurnMinProto : turn to protobuf with certain fields
func (c *IntactClass) TurnMinProto() *protocol.Class {
	return &protocol.Class{
		Tutor:      c.TutorName,
		Id:         c.Class.Id,
		Name:       c.Class.Name,
		Number:     int64(c.Class.Number),
		IsCheck:    c.Class.IsCheck == 1,
		CreateTime: c.Class.CreateTime,
	}
}

// TurnProtoToIntactClass : turn protobuf to IntactClass
func TurnProtoToIntactClass(p *protocol.Class) *IntactClass {
	c := &model.Class{
		Id:           p.Id,
		CreateTime:   time.Now().Unix(),
		Introduction: p.Introduction,
		Name:         p.Name,
	}
	if p.IsCheck {
		c.IsCheck = 1
	}
	c.IsCheck = 0

	// set problem test data
	var announcements []*AnnouncementWithName
	for _, announcement := range p.Announcements {
		a := model.Announcement{
			ClassId:    p.Id,
			Detail:     announcement.Detail,
			CreateTime: time.Now().Unix(),
		}
		announcements = append(announcements, &AnnouncementWithName{
			Announcement: a,
		})
	}

	return &IntactClass{
		Class:         c,
		Announcements: announcements,
	}
}
