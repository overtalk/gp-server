package user

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetUsers : get users
func (u *User) GetUsers(r *http.Request) proto.Message {
	req := &protocol.GetUsersReq{}
	resp := &protocol.GetUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	var users []*model.User
	var err error
	// get required users informations
	role := req.Role
	if req.GetAll {
		users, err = u.db.GetUsers(req.PageNum, req.PageIndex)
	} else {
		users, err = u.db.GetUsersByRole(req.PageNum, req.PageIndex, int64(role))
	}
	if err != nil {
		logger.Sugar.Errorf("failed to get users : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get users"
		return resp
	}

	// add users informations to response
	for _, user := range users {
		resp.Users = append(resp.Users, transform.UserToProto(user))
	}

	// get all number
	usersNum, err := u.db.GetUsersNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of users : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of users"
		return resp
	}
	resp.Total = usersNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}

// AddUsers : add users to db
func (u *User) AddUsers(r *http.Request) proto.Message {
	req := &protocol.AddUsersReq{}
	resp := &protocol.AddUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// add users
	ts := time.Now().Unix()
	for _, protoUser := range req.Users {
		user := transform.ProtoToUser(protoUser)
		user.LastLogin = ts
		user.Create = ts
		if err := u.db.AddUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, transform.UserToProto(user))
		}
	}

	return resp
}

// UpdateUsers : update users
func (u *User) UpdateUsers(r *http.Request) proto.Message {
	req := &protocol.UpdateUsersReq{}
	resp := &protocol.UpdateUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// update users
	for _, protoUser := range req.Users {
		user := transform.ProtoToUser(protoUser)
		if err := u.db.UpdateUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, transform.UserToProto(user))
		}
	}

	return resp
}

// DelUsers : delete users
func (u *User) DelUsers(r *http.Request) proto.Message {
	req := &protocol.DelUsersReq{}
	resp := &protocol.DelUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// delete users
	for _, userID := range req.UsersId {
		if err := u.db.DelUser(userID); err != nil {
			resp.Fail = append(resp.Fail, userID)
		} else {
			resp.Succeed = append(resp.Succeed, userID)
		}
	}

	return resp
}
