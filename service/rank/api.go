package rank

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetRankList : get rank list info
func (k *Rank) GetRankList(r *http.Request) proto.Message {
	req := &protocol.RankListReq{}
	resp := &protocol.RankListResp{Status: &protocol.Status{}}
	// get token and data
	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		logger.Sugar.Error(err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = err.Error()
		return resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	// check token
	_, err = k.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	rankItems := k.getRanksFromCache()
	for _, item := range rankItems {
		fmt.Println(item)
	}
	total := int64(len(rankItems))
	start := (req.PageIndex - 1) * req.PageNum
	end := start + req.PageNum
	if total > start {
		if total >= end {
			resp.RankItems = rankItems[start:end]
		} else {
			resp.RankItems = rankItems[start:total]
		}
	}
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	resp.Total = total
	return resp
}
