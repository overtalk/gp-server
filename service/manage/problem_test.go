package manage_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/manage"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetProblems(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	reqBytes, err := proto.Marshal(&protocol.GetProblemsReq{})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := managerModule.GetProblems(args)
	resp := data.(*protocol.GetProblemsResp)
	if resp.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Code)])
		return
	}

	for _, problem := range resp.Problems {
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
	managerModule := manage.NewBackStageManager(dataStorage)

	reqBytes, err := proto.Marshal(&protocol.GetProblemByIDReq{
		Id: 1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := managerModule.GetProblemByID(args)
	resp := data.(*protocol.GetProblemByIDResp)
	if resp.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Code)])
		return
	}

	t.Logf("%v", resp.Problem)
	t.Logf("%v", resp.Problem.InOutExamples)
	t.Logf("%v", resp.Problem.JudgeLimit)
	t.Logf("%v", resp.Problem.Tags)
}

func TestUserManage_AddProblem(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	reqBytes, err := proto.Marshal(&protocol.AddProblemReq{
		Problem: &protocol.Problem{
			Title:       "这是一个测试题目(求和)",
			Description: "题目描述",
			In:          "输入描述",
			Out:         "输出描述",
			Hint:        "没有提示",
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
			JudgeLimit: &protocol.ProblemJudgeLimit{
				Time: "62s",
				Mem:  "7m",
			},
			Tags:       []string{"求和", "数组", "树"},
			Difficluty: protocol.ProblemDifficluty_HARD,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := managerModule.AddProblem(args)
	resp := data.(*protocol.AddProblemResp)
	if resp.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Code)])
		return
	}
}
