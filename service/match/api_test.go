package match_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/match"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_NewMatch(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.NewMatchReq{
		Paper: &protocol.Paper{
			Difficulty:      1,
			ProblemNum:      3,
			KnowledgePoints: []int64{1, 2, 3},
		},
		Match: &protocol.Match{
			IsPublic:     false,
			StartTime:    time.Now().Unix(),
			Duration:     10000,
			Name:         "呵呵呵考试",
			Intriduction: "asdfasdf",
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.NewMatch(r)
	resp := data.(*protocol.NewMatchResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}
}

func TestUserManage_GetMatches(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetMatchesReq{
		PageIndex: 1,
		PageNum:   3,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetMatches(r)
	resp := data.(*protocol.GetMatchesResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Logf("index = %d", resp.PageIndex)
	t.Logf("num = %d", resp.PageNum)
	t.Logf("total = %d", resp.Total)
	for _, match := range resp.Matches {
		t.Log(match)
	}
}

func TestUserManage_GetMatchByID(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetMatchByIDReq{
		Id: 1,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetMatchByID(r)
	resp := data.(*protocol.GetMatchByIDResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.Match)
}

func TestUserManage_GetPaperByID(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.GetPaperByIDReq{
		Id: 1,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.GetPaperByID(r)
	resp := data.(*protocol.GetPaperByIDResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	t.Log(resp.Paper.KnowledgePoints)
	t.Log("题目数量", len(resp.Paper.Problems))
	t.Log(resp.Paper.Problems)
}

func TestUserManage_EditMatch(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := match.NewMatch(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.EditMatchReq{
		Match: &protocol.Match{
			Id:   1,
			Name: "asdfdsfgsdfgs",
		},
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.EditMatch(r)
	resp := data.(*protocol.EditMatchResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}
}
