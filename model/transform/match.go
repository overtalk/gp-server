package transform

import (
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
