package user_manage

import (
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetUsers : get users
func (m *UserManager) GetUsers(args map[string]interface{}) interface{} {
	resp := &protocol.GetUsersResp{}
	if err := utils.CheckArgs(args, module.Request, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	req := &protocol.GetUsersReq{}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	userID, err := m.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		logger.Sugar.Errorf("invaild token : %v", err)
		resp.Code = protocol.Code_INVAILD_TOKEN
		return resp
	}

	user, err := m.db.GetUserByID(userID)
	if err != nil {
		logger.Sugar.Errorf("failed to get user by id[%d] : %v", userID, err)
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	if user.OperationAuth == int(protocol.Role_MANAGER) {
		logger.Sugar.Errorf("failed to get users[permission denied] for user[id = %d, role = %s]", userID, protocol.Role_name[int32(user.Role)])
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	role := req.Role
	if req.GetAll {
		role = -1
	}
	users, err := m.db.GetUsersByRole(int64(role))
	if err != nil {
		logger.Sugar.Errorf("failed to get users : %v", userID, err)
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	for _, user := range users {
		resp.Users = append(resp.Users, user.TurnProto())
	}

	return resp
}
