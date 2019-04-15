package class

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetClasses : get all classes
func (c *Class) GetClasses(r *http.Request) proto.Message {
	req := &protocol.GetClassesReq{}
	resp := &protocol.GetClassesResp{Status: &protocol.Status{}}

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

	classes, err := c.db.GetClasses(req.PageNum, req.PageIndex)
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
	classesNum, err := c.db.GetClassNum()
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

// GetClassByID : get a certain class
func (c *Class) GetClassByID(r *http.Request) proto.Message {
	req := &protocol.GetClassByIDReq{}
	resp := &protocol.GetClassByIDResp{Status: &protocol.Status{}}

	_, status := c.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	class, err := c.db.GetClassByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get the class : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to get the class")
		return resp
	}

	resp.Class = class.TurnProto()
	return resp
}

// AddClass : add a new class
func (c *Class) AddClass(r *http.Request) proto.Message {
	req := &protocol.AddClassReq{}
	resp := &protocol.AddClassResp{Status: &protocol.Status{}}

	user, status := c.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// set publisher and tutor
	intactClass := transform.TurnProtoToIntactClass(req.Class)
	intactClass.Class.Tutor = user.Id
	for _, announcement := range intactClass.Announcements {
		announcement.Publisher = user.Id
	}
	if err := c.db.AddClass(intactClass); err != nil {
		logger.Sugar.Errorf("failed to add the class : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to add the class")
		return resp
	}

	resp.IsSuccess = true
	return resp
}

// EditClass : update the message of Class
func (c *Class) EditClass(r *http.Request) proto.Message {
	req := &protocol.EditClassReq{}
	resp := &protocol.EditClassResp{Status: &protocol.Status{}}

	_, status := c.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	intactClass := transform.TurnProtoToIntactClass(req.Class)
	for _, announcement := range intactClass.Announcements {
		announcement.ClassId = intactClass.Class.Id
	}
	if err := c.db.UpdateClass(intactClass); err != nil {
		logger.Sugar.Errorf("failed to edit the class : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to edit the class")
		return resp
	}

	resp.IsSuccess = true
	return resp
}

// MemberManage : member manage
func (c *Class) MemberManage(r *http.Request) proto.Message {
	req := &protocol.MemberManageReq{}
	resp := &protocol.MemberManageResp{Status: &protocol.Status{}}

	_, status := c.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	if err := c.db.MemberManage(int64(req.ManageType), req.ClassId, req.MemberId); err != nil {
		logger.Sugar.Errorf("failed to edit the member of class : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to edit the member of class")
		return resp
	}

	resp.IsSuccess = true
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

	members, num, err := c.db.GetMembers(req.ClassId, req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get class members : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get class members"
		return resp
	}

	for _, member := range members {
		resp.Menbers = append(resp.Menbers, transform.UserClassToProto(member))
	}
	resp.PageNum = req.PageNum
	resp.PageIndex = req.PageIndex
	resp.Total = num

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
