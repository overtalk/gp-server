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
		Duration:     int(match.Duration),
		Introduction: match.Intriduction,
		IsPublic:     isPublic,
		StartTime:    match.StartTime,
		Title:        match.Name,
	}
}

// MatchToProto : turn modle Match to proto
func MatchToProto(u *model.Match) *protocol.Match {
	isOver := false
	if u.StartTime+int64(u.Duration) < time.Now().Unix() {
		isOver = true
	}
	return &protocol.Match{
		Id:           u.Id,
		IsPublic:     u.IsPublic == 1,
		StartTime:    u.StartTime,
		Duration:     int64(u.Duration),
		Name:         u.Title,
		Intriduction: u.Introduction,
		IsOver:       isOver,
	}
}
