package problem_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/problem"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetProblems(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := problem.NewProblem(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetProblemsReq{
		Tag:       3,
		PageIndex: 1,
		PageNum:   5,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetProblems(r)
	resp := data.(*protocol.GetProblemsResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	for _, problem := range resp.Problems {
		t.Logf("%+v", problem.Publisher)
		t.Logf("%+v", problem)
	}
}

func TestUserManage_GetProblemByID(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := problem.NewProblem(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetProblemByIDReq{Id: 1})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetProblemByID(r)
	resp := data.(*protocol.GetProblemByIDResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("%v", resp.Problem)
	t.Logf("%v", resp.Problem.InOutExamples)
	t.Logf("%v", resp.Problem.JudgeLimitMem)
	t.Logf("%v", resp.Problem.JudgeLimitTime)
	t.Logf("%v", resp.Problem.Tags)
}

func TestUserManage_AddProblem(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := problem.NewProblem(dataStorage)

	req := &protocol.AddProblemReq{
		Problem: &protocol.Problem{
			Title:          "这是一个测试题目(求和)",
			Description:    "题目描述",
			In:             "输入描述",
			Out:            "输出描述",
			Hint:           "没有提示",
			JudgeLimitMem:  5,
			JudgeLimitTime: 3,
			JudgeFile:      "1120564623934791680",
			Tags:           []int64{1, 3},
			Difficulty:     1,
			InOutExamples: []*protocol.ProblemExample{
				&protocol.ProblemExample{
					Input:  "1 1",
					Output: "2",
				},
				&protocol.ProblemExample{
					Input:  "2 2",
					Output: "4",
				},
			},
		},
	}
	r, err := utils.MockHTTPReq("POST", "1", req)
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.AddProblem(r)
	resp := data.(*protocol.AddProblemResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}
}

func TestUserManage_EditProblem(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := problem.NewProblem(dataStorage)

	req := &protocol.EditProblemReq{
		Problem: &protocol.Problem{
			Id:   6,
			Hint: "在TestUserManage_EditProblem中被修改",
		},
	}
	r, err := utils.MockHTTPReq("POST", "1", req)
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.EditProblem(r)
	resp := data.(*protocol.EditProblemResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}
}
