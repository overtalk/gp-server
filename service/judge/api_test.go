package judge_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/judge"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestJudge_judge(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	module := judge.NewJudge(dataStorage)

	r, err := utils.MockHTTPReq("POST", "1", &protocol.JudgeRequest{
		Id: 1,
		Src: `
		#include <iostream>
		
		using namespace std;
		
		int main()
		{
			int a,b;
			cin >> a >> b;
			cout << a+b << endl;
			return 0;
		}
		`,
		Language: protocol.Language_C_PLUS,
	})
	if err != nil {
		t.Errorf("failed to mock http request : %v", err)
		return
	}

	data := module.Judge(r)
	resp := data.(*protocol.JudgeResponse)
	if resp.Status.Code != protocol.Code_OK {
		t.Errorf("resp.Code[%s] != protocol.Code_OK", protocol.Code_name[int32(resp.Status.Code)])
		return
	}

	for _, result := range resp.Results {
		t.Log(result)
	}
}
