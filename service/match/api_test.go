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
