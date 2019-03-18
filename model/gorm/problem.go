package model

import (
	"encoding/json"

	"github.com/qinhan-shu/gp-server/protocol"
)

// Problem : teble `problem`
type Problem struct {
	ID             int64  `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Title          string `gorm:"type : varchar(300); not null"`
	Description    string `gorm:"type : text; not null"`
	InDescription  string `gorm:"type : text; not null"`
	OutDescription string `gorm:"type : text; not null"`
	Hint           string `gorm:"type : text"`
	Example        string `gorm:"type : text; not null"`
	JudgeFile      string `gorm:"type : varchar(100); not null"`
	JudgeLimit     string `gorm:"type : json; not null"`
	Tags           string `gorm:"type : varchar(300); not null"`
	Difficulty     int    `gorm:"type : tinyint(4); not null; default : 0"`
	LastUsed       int64  `gorm:"type : int(64); not null"`
	UsedTime       int64  `gorm:"type : int(20); not null; default : 0"`
	SubmitTime     int64  `gorm:"type : int(64); not null; default : 0"` // 提交总次数
	Ac             int64  `gorm:"type : int(64); not null; default : 0"` // 通过次数
	Wa             int64  `gorm:"type : int(64); not null; default : 0"` // 答案错误次数
	Tle            int64  `gorm:"type : int(64); not null; default : 0"` // 超时次数
	Mle            int64  `gorm:"type : int(64); not null; default : 0"` // 超内存次数
	Pe             int64  `gorm:"type : int(64); not null; default : 0"` // 格式错误次数
	Ce             int64  `gorm:"type : int(64); not null; default : 0"` // 编译次数
}

// TurnProto : turn Problem to protobuf
func (p *Problem) TurnProto() *protocol.Problem {
	problemProtobuf := &protocol.Problem{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		In:          p.InDescription,
		Out:         p.OutDescription,
		Hint:        p.Hint,
		Difficluty:  protocol.ProblemDifficluty(p.Difficulty),
		SubmitTime:  p.SubmitTime,
		AcceptTime:  p.Ac,
	}
	json.Unmarshal([]byte(p.Example), problemProtobuf.InOutExamples)
	json.Unmarshal([]byte(p.Tags), problemProtobuf.Tags)
	json.Unmarshal([]byte(p.JudgeLimit), problemProtobuf.JudgeLimit)
	return problemProtobuf
}

// TurnMinProto : turn to protobuf with certain fields
func (p *Problem) TurnMinProto() *protocol.Problem {
	return &protocol.Problem{
		Id:         p.ID,
		Title:      p.Title,
		Difficluty: protocol.ProblemDifficluty(p.Difficulty),
		SubmitTime: p.SubmitTime,
		AcceptTime: p.Ac,
	}
}

// IsInited : check the default value of each fields
func (p *Problem) IsInited() bool {
	return p.Title != "" && p.Description != "" && p.InDescription != "" && p.OutDescription != "" && p.Example != "" && p.JudgeLimit != "" && p.JudgeFile != "" && p.Tags != ""
}

// TurnProblem : turn protobuf to problem
func TurnProblem(p *protocol.Problem) *Problem {
	inOutExamples, _ := json.Marshal(p.InOutExamples)
	judgeLimit, _ := json.Marshal(p.JudgeLimit)
	tags, _ := json.Marshal(p.Tags)
	return &Problem{
		Title:          p.Title,
		Description:    p.Description,
		InDescription:  p.In,
		OutDescription: p.Out,
		Hint:           p.Hint,
		Example:        string(inOutExamples),
		JudgeLimit:     string(judgeLimit),
		Tags:           string(tags),
		// TODO:
		JudgeFile: "/file",
	}
}
