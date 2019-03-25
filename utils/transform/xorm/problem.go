package transform

import (
	model_utils "github.com/qinhan-shu/gp-server/model"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// ProblemToProto : turn Problem to protobuf
func ProblemToProto(p *model_utils.IntactProblem) *protocol.Problem {
	problemProtobuf := &protocol.Problem{
		Id:             p.Detail.Id,
		Title:          p.Detail.Title,
		Description:    p.Detail.Description,
		In:             p.Detail.InDescription,
		Out:            p.Detail.OutDescription,
		Hint:           p.Detail.Hint,
		Difficluty:     protocol.ProblemDifficluty(p.Detail.Difficulty),
		SubmitTime:     int64(p.Detail.SubmitTime),
		AcceptTime:     int64(p.Detail.Ac),
		JudgeLimitMem:  int64(p.Detail.JudgeLimitMem),
		JudgeLimitTime: int64(p.Detail.JudgeLimitTime),
	}
	// set in and out example
	var example []*protocol.ProblemExample
	for _, testData := range p.InAndOutExample {
		example = append(example, &protocol.ProblemExample{
			Input:  testData.In,
			Output: testData.Out,
		})
	}
	problemProtobuf.InOutExamples = example

	// set tags
	var tags []int64
	for _, tag := range p.Tags {
		tags = append(tags, int64(tag.TagId))
	}
	problemProtobuf.Tags = tags
	return problemProtobuf
}

// ProblemToMinProto : turn to protobuf with certain fields
func ProblemToMinProto(p *model_utils.IntactProblem) *protocol.Problem {
	return &protocol.Problem{
		Id:         p.Detail.Id,
		Title:      p.Detail.Title,
		Difficluty: protocol.ProblemDifficluty(p.Detail.Difficulty),
		SubmitTime: int64(p.Detail.SubmitTime),
		AcceptTime: int64(p.Detail.Ac),
	}
}

// IsInited : check the default value of each fields
func IsInited(p *model.Problem) bool {
	return p.Title != "" && p.Description != "" && p.InDescription != "" && p.OutDescription != ""
}

// ProtoToProblem : turn protobuf to problem
func ProtoToProblem(p *protocol.Problem) *model_utils.IntactProblem {
	problem := &model.Problem{
		Id:             p.Id,
		Title:          p.Title,
		Description:    p.Description,
		InDescription:  p.In,
		OutDescription: p.Out,
		Hint:           p.Hint,
		JudgeLimitMem:  int(p.JudgeLimitMem),
		JudgeLimitTime: int(p.JudgeLimitTime),
		Difficulty:     int(p.Difficluty),
	}
	// set problem test data
	var inAndOutExample []*model.TestData
	for _, testData := range p.InOutExamples {
		inAndOutExample = append(inAndOutExample, &model.TestData{
			ProblemId: problem.Id,
			In:        testData.Input,
			Out:       testData.Output,
		})
	}
	// set tags
	var tags []*model.ProblemTag
	for _, tag := range p.Tags {
		tags = append(tags, &model.ProblemTag{
			ProblemId: p.Id,
			TagId:     int(tag),
		})
	}
	return &model_utils.IntactProblem{
		Detail:          problem,
		InAndOutExample: inAndOutExample,
		Tags:            tags,
	}
}
