package class

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// SearchClass : get classes by keyword
func (c *Class) SearchClass(r *http.Request) proto.Message {
	req := &protocol.SearchClassesReq{}
	resp := &protocol.SearchClassesResp{Status: &protocol.Status{}}

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
	_, err = c.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	classes, err := c.db.GetClasses(req.PageNum, req.PageIndex, req.Keyword)
	if err != nil {
		logger.Sugar.Errorf("failed to get classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get classes"
		return resp
	}

	for _, class := range classes {
		resp.Classes = append(resp.Classes, class.TurnMinProto())
	}

	// get all number
	classesNum, err := c.db.GetClassNum(req.Keyword)
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of classes"
		return resp
	}

	resp.Total = classesNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}

// GetMembers : get all members
func (c *Class) GetMembers(r *http.Request) proto.Message {
	req := &protocol.GetMemberReq{}
	resp := &protocol.GetMemberResp{Status: &protocol.Status{}}
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
	_, err = c.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	members, err := c.db.GetMembers(req.ClassId)
	if err != nil {
		logger.Sugar.Errorf("failed to get class members : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get class members"
		return resp
	}

	for _, member := range members {
		resp.Members = append(resp.Members, transform.UserClassToProto(member))
	}

	return resp
}

// EnterClass : enter class
func (c *Class) EnterClass(r *http.Request) proto.Message {
	req := &protocol.EnterClassReq{}
	resp := &protocol.EnterClassResp{Status: &protocol.Status{}}
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
	userID, err := c.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	if err := c.db.EnterClass(userID, req.ClassId); err != nil {
		logger.Sugar.Errorf("failed to enter class : %v", err)
		resp.IsSuccess = false
	} else {
		resp.IsSuccess = true
	}

	return resp
}

// QuitClass : quit class
func (c *Class) QuitClass(r *http.Request) proto.Message {
	req := &protocol.QuitClassReq{}
	resp := &protocol.QuitClassResp{Status: &protocol.Status{}}
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
	userID, err := c.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	c.db.QuitClass(userID, req.ClassId)
	resp.IsSuccess = true

	return resp
}

// ApplyEnterRequest : teacher agree or disagree the enter class request of student
func (c *Class) ApplyEnterRequest(r *http.Request) proto.Message {
	req := &protocol.ApplyEnterRequestReq{}
	resp := &protocol.ApplyEnterRequestResp{Status: &protocol.Status{}}
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
	_, err = c.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	if err := c.db.ApplyEnterRequest(req.UserId, req.ClassId, req.IsApply); err != nil {
		logger.Sugar.Errorf("failed to apply enter class request : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}
	resp.IsSuccess = true

	return resp
}
