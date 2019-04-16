package transform

import (
	"encoding/json"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// IntactProblem : intact problem
type IntactProblem struct {
	model.Problem   `xorm:"extends"`
	InAndOutExample []*model.TestData
	Name            string
}

func (IntactProblem) TableName() string {
	return "problem"
}

// TurnProto : turn Problem to protobuf
func (p *IntactProblem) TurnProto() *protocol.Problem {
	problemProtobuf := &protocol.Problem{
		Id:             p.Id,
		Title:          p.Title,
		Description:    p.Description,
		In:             p.InDescription,
		Out:            p.OutDescription,
		Hint:           p.Hint,
		Difficulty:     int64(p.Difficulty),
		SubmitTime:     int64(p.SubmitTime),
		AcceptTime:     int64(p.Ac),
		JudgeLimitMem:  int64(p.JudgeLimitMem),
		JudgeLimitTime: int64(p.JudgeLimitTime),
		Publisher:      p.Name,
		CreateTime:     p.CreateTime,
	}
	// set in and out example
	var example []*protocol.ProblemExample
	for _, testData := range p.InAndOutExample {
		example = append(example, &protocol.ProblemExample{
			Index:  testData.Index,
			Input:  testData.In,
			Output: testData.Out,
		})
	}
	problemProtobuf.InOutExamples = example

	// set tags
	tags := new([]int64)
	json.Unmarshal([]byte(p.Tags), tags)

	problemProtobuf.Tags = *tags
	return problemProtobuf
}

// TurnMinProto : turn to protobuf with certain fields
func (p *IntactProblem) TurnMinProto() *protocol.Problem {
	return &protocol.Problem{
		Id:         p.Id,
		Title:      p.Title,
		Difficulty: int64(p.Difficulty),
		SubmitTime: int64(p.SubmitTime),
		AcceptTime: int64(p.Ac),
		Publisher:  p.Name,
		CreateTime: p.CreateTime,
	}
}

// IsInited : check the default value of each fields
func IsInited(p *model.Problem) bool {
	return p.Title != "" && p.Description != "" && p.InDescription != "" && p.OutDescription != ""
}

// ProtoToProblem : turn protobuf to problem
func ProtoToProblem(p *protocol.Problem) *IntactProblem {
	problem := model.Problem{
		Id:             p.Id,
		Title:          p.Title,
		Description:    p.Description,
		InDescription:  p.In,
		OutDescription: p.Out,
		Hint:           p.Hint,
		JudgeLimitMem:  int(p.JudgeLimitMem),
		JudgeLimitTime: int(p.JudgeLimitTime),
		Difficulty:     int(p.Difficulty),
	}
	// set problem test data
	var inAndOutExample []*model.TestData
	for _, testData := range p.InOutExamples {
		inAndOutExample = append(inAndOutExample, &model.TestData{
			ProblemId: problem.Id,
			Index:     testData.Index,
			In:        testData.Input,
			Out:       testData.Output,
		})
	}
	// set tags
	tags, _ := json.Marshal(p.Tags)
	problem.Tags = string(tags)

	return &IntactProblem{
		Problem:         problem,
		InAndOutExample: inAndOutExample,
	}
}
