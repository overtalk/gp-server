package manage

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/gorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetUsers : get users
func (m *BackStageManage) GetUsers(c *gin.Context) (int, interface{}) {
	code := http.StatusOK
	req := &protocol.GetUsersReq{}
	resp := &protocol.GetUsersResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to get users[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	// get required users informations
	role := req.Role
	if req.GetAll {
		role = -1
	}
	users, err := m.db.GetUsersByRole(int64(role))
	if err != nil {
		logger.Sugar.Errorf("failed to get users : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}

	// add users informations to response
	for _, user := range users {
		resp.Users = append(resp.Users, user.TurnProto())
	}

	return code, resp
}

// AddUsers : add users to db
func (m *BackStageManage) AddUsers(c *gin.Context) (int, interface{}) {
	code := http.StatusOK
	req := &protocol.AddUsersReq{}
	resp := &protocol.AddUsersResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to add users[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	// add users
	ts := time.Now().Unix()
	for _, protoUser := range req.Users {
		user := model.TurnUser(protoUser)
		user.LastLogin = ts
		user.Create = ts
		if err := m.db.AddUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, user.TurnProto())
		}
	}

	return code, resp
}

// UpdateUsers : update users
func (m *BackStageManage) UpdateUsers(c *gin.Context) (int, interface{}) {
	code := http.StatusOK
	req := &protocol.UpdateUsersReq{}
	resp := &protocol.UpdateUsersResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to update users[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	// update users
	for _, protoUser := range req.Users {
		user := model.TurnUser(protoUser)
		if err := m.db.UpdateUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, user.TurnProto())
		}
	}

	return code, resp
}

// DelUsers : delete users
func (m *BackStageManage) DelUsers(c *gin.Context) (int, interface{}) {
	code := http.StatusOK
	req := &protocol.DelUsersReq{}
	resp := &protocol.DelUsersResp{}
	// get token and data
	data, token, err := getReqAndToken(c)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		code = http.StatusUnauthorized
		return code, resp
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		code = http.StatusInternalServerError
		return code, resp
	}
	if !authCheck(user.Role) {
		logger.Sugar.Errorf("failed to delete users[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		code = http.StatusUnauthorized
		return code, resp
	}

	// delete users
	for _, userID := range req.UsersId {
		if err := m.db.DelUser(userID); err != nil {
			resp.Fail = append(resp.Fail, userID)
		} else {
			resp.Succeed = append(resp.Succeed, userID)
		}
	}

	return code, resp
}
