package conf

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetTags : get all tags
func (a *Conf) GetTags(r *http.Request) proto.Message {
	req := &protocol.GetTagsReq{}
	resp := &protocol.GetTagsResp{Status: &protocol.Status{}}
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
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	tags, err := a.db.GetAllTagsByPage(req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get tags by page : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "failed to get tags by page"
		return resp
	}

	t := make(map[int64]string)
	for _, v := range tags {
		t[int64(v.Id)] = v.Detail
	}
	resp.Tags = t

	tags, err = a.db.GetAllTag()
	if err != nil {
		logger.Sugar.Errorf("failed to get tags by page : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "failed to get tags by page"
		return resp
	}

	resp.Total = int64(len(tags))
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum

	return resp
}

// AddTag : add tag
func (a *Conf) AddTag(r *http.Request) proto.Message {
	req := &protocol.AddTagReq{}
	resp := &protocol.AddTagResp{Status: &protocol.Status{}}
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
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	tag := &model.Tag{
		Detail: req.Tag,
	}

	if err := a.db.AddTag(tag); err != nil {
		logger.Sugar.Errorf("failed to add tag : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "failed to add tag"
		return resp
	}

	resp.IsSuccess = true
	return resp
}

// UpdateTag : update tag
func (a *Conf) UpdateTag(r *http.Request) proto.Message {
	req := &protocol.UpdateTagReq{}
	resp := &protocol.UpdateTagResp{Status: &protocol.Status{}}
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
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	tag := &model.Tag{
		Id:     int(req.Id),
		Detail: req.Tag,
	}

	if err := a.db.UpdateTag(tag); err != nil {
		logger.Sugar.Errorf("failed to update tag : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "failed to update tag"
		return resp
	}

	resp.IsSuccess = true
	return resp
}
