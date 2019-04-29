package transform

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// UserProblemToProto : turn user problem to protobuf
func UserProblemToProto(u *model.UserProblem) *protocol.SubmitRecord {
	return &protocol.SubmitRecord{
		ProblemId:   u.ProblemId,
		UserId:      u.UserId,
		SubmitTime:  int64(u.SubmitTime),
		IsPass:      u.Ispass == 1,
		RunningMem:  int64(u.RunningMem),
		RunningTime: int64(u.RunningTime),
		Code:        u.Code,
		RunningLan:  int64(u.RunningLanguage),
	}
}

// UserProblemWithDifficulty : user problem with difficulty
type UserProblemWithDifficulty struct {
	model.UserProblem `xorm:"extends"`
	Difficulty        int
}

func (UserProblemWithDifficulty) TableName() string {
	return "user_problem"
}

type UserProblemWithTags struct {
	model.UserProblem `xorm:"extends"`
	Tags              string
}

func (UserProblemWithTags) TableName() string {
	return "user_problem"
}
