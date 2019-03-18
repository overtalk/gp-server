package model

import (
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
	JudgeLimit     string `gorm:"type : json"`
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
	return &protocol.Problem{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		In:          p.InDescription,
		Out:         p.OutDescription,
		Hint:        p.Hint,
		Difficluty:  protocol.ProblemDifficluty(p.Difficulty),
		SubmitTime:  p.SubmitTime,
		AcceptTime:  p.Ac,
		// TODO:
		InOutExamples: nil,
		Tags:          nil,
		JudgeLimit: &protocol.ProblemJudgeLimit{
			Time: "",
			Mem:  "",
		},
	}
}
