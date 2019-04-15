package transform

import (
	"time"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// ProtoToMatch : turn protobuf to match
func ProtoToMatch(match *protocol.Match) *model.Match {
	isPublic := 0
	if match.IsPublic {
		isPublic = 1
	}
	return &model.Match{
		Id:           match.Id,
		EndTime:      match.EndTime,
		Introduction: match.Introduction,
		IsPublic:     isPublic,
		StartTime:    match.StartTime,
		Title:        match.Name,
	}
}

// MatchToMinProto : turn modle Match to min proto
func MatchToMinProto(u *model.Match) *protocol.Match {
	isOver := false
	if u.EndTime < time.Now().Unix() {
		isOver = true
	}
	return &protocol.Match{
		Id:        u.Id,
		IsPublic:  u.IsPublic == 1,
		StartTime: u.StartTime,
		EndTime:   u.EndTime,
		Name:      u.Title,
		IsOver:    isOver,
		PaperId:   u.PaperId,
	}
}

// MatchToProto : turn modle Match to proto
func MatchToProto(u *model.Match) *protocol.Match {
	isOver := false
	if u.EndTime < time.Now().Unix() {
		isOver = true
	}
	return &protocol.Match{
		Id:           u.Id,
		IsPublic:     u.IsPublic == 1,
		StartTime:    u.StartTime,
		EndTime:      u.EndTime,
		Name:         u.Title,
		Introduction: u.Introduction,
		IsOver:       isOver,
		PaperId:      u.PaperId,
	}
}
