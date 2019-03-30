package rank_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/rank"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestRank_GetRankList(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	rankModule := rank.NewRank(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.RankListReq{
		PageIndex: 3,
		PageNum:   4,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := rankModule.GetRankList(r)
	resp := data.(*protocol.RankListResp)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	for _, item := range resp.RankItems {
		t.Log(item)
	}
}
