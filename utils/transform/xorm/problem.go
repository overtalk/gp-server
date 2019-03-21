package turn

import (
	"encoding/json"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// ProblemToProto : turn Problem to protobuf
func ProblemToProto(p *model.Problem) *protocol.Problem {
	problemProtobuf := &protocol.Problem{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		In:          p.InDescription,
		Out:         p.OutDescription,
		Hint:        p.Hint,
		Difficluty:  protocol.ProblemDifficluty(p.Difficulty),
		SubmitTime:  int64(p.SubmitTime),
		AcceptTime:  int64(p.Ac),
	}
	json.Unmarshal([]byte(p.Example), &problemProtobuf.InOutExamples)
	json.Unmarshal([]byte(p.Tags), &problemProtobuf.Tags)
	json.Unmarshal([]byte(p.JudgeLimit), &problemProtobuf.JudgeLimit)
	return problemProtobuf
}

// ProblemToMinProto : turn to protobuf with certain fields
func ProblemToMinProto(p *model.Problem) *protocol.Problem {
	return &protocol.Problem{
		Id:         p.Id,
		Title:      p.Title,
		Difficluty: protocol.ProblemDifficluty(p.Difficulty),
		SubmitTime: int64(p.SubmitTime),
		AcceptTime: int64(p.Ac),
	}
}

// IsInited : check the default value of each fields
func IsInited(p *model.Problem) bool {
	return p.Title != "" && p.Description != "" && p.InDescription != "" && p.OutDescription != "" && p.Example != "" && p.JudgeLimit != "" && p.JudgeFile != "" && p.Tags != ""
}

// ProtoToProblem : turn protobuf to problem
func ProtoToProblem(p *protocol.Problem) *model.Problem {
	problem := &model.Problem{
		Id:             p.Id,
		Title:          p.Title,
		Description:    p.Description,
		InDescription:  p.In,
		OutDescription: p.Out,
		Hint:           p.Hint,
	}
	if p.InOutExamples != nil {
		inOutExamples, _ := json.Marshal(p.InOutExamples)
		problem.Example = string(inOutExamples)
	}
	if p.JudgeLimit != nil {
		judgeLimit, _ := json.Marshal(p.JudgeLimit)
		problem.JudgeLimit = string(judgeLimit)
	}
	if p.Tags != nil {
		tags, _ := json.Marshal(p.Tags)
		problem.Tags = string(tags)
	}
	return problem
}
