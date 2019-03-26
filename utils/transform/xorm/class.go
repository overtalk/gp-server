package transform

import (
	"time"

	model_utils "github.com/qinhan-shu/gp-server/model"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// TurnIntactClassToProto : turn intactProblem to protobuf
func TurnIntactClassToProto(intactClass *model_utils.IntactClass) *protocol.Class {
	protobufClass := &protocol.Class{
		Tutor:        intactClass.TutorName,
		Id:           intactClass.Class.Id,
		Name:         intactClass.Class.Name,
		Introduction: intactClass.Class.Introduction,
		Number:       int64(intactClass.Class.Number),
		IsCheck:      intactClass.Class.IsCheck == 1,
		CreateTime:   intactClass.Class.CreateTime,
	}
	for _, announcement := range intactClass.Announcements {
		protobufClass.Announcements = append(protobufClass.Announcements, &protocol.Announcement{
			Publisher:  intactClass.TutorName,
			Detail:     announcement.Detail,
			CreateTime: announcement.CreateTime,
		})
	}
	return protobufClass
}

// TurnIntactClassToMinProto : turn to protobuf with certain fields
func TurnIntactClassToMinProto(intactClass *model_utils.IntactClass) *protocol.Class {
	return &protocol.Class{
		Tutor:      intactClass.TutorName,
		Id:         intactClass.Class.Id,
		Name:       intactClass.Class.Name,
		Number:     int64(intactClass.Class.Number),
		IsCheck:    intactClass.Class.IsCheck == 1,
		CreateTime: intactClass.Class.CreateTime,
	}
}

// TurnProtoToIntactClass : turn protobuf to IntactClass
func TurnProtoToIntactClass(p *protocol.Class) *model_utils.IntactClass {
	c := &model.Class{
		CreateTime:   p.CreateTime,
		Id:           p.Id,
		Introduction: p.Introduction,
		Name:         p.Name,
	}
	if p.IsCheck {
		c.IsCheck = 1
	}
	c.IsCheck = 0

	// set problem test data
	var announcements []*model.Announcement
	for _, announcement := range p.Announcements {
		announcements = append(announcements, &model.Announcement{
			ClassId:    p.Id,
			Detail:     announcement.Detail,
			CreateTime: time.Now().Unix(),
		})
	}

	return &model_utils.IntactClass{
		Class:         c,
		Announcements: announcements,
	}
}
