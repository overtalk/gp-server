package transform

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// UserToProto : turn user to protobuf
// account & password is not allowed to send to client
func UserToProto(u *model.User) *protocol.UserInfo {
	return &protocol.UserInfo{
		Account:   u.Account,
		Create:    u.Create,
		Email:     u.Email,
		Id:        u.Id,
		LastLogin: u.LastLogin,
		Name:      u.Name,
		// Password:  u.Password,
		Phone:  u.Phone,
		Role:   int64(u.Role),
		School: u.School,
		Sex:    u.Sex == 1,
	}
}

// IsUserInited : check the default value of each fields
func IsUserInited(u *model.User) bool {
	return u.Account != "" && u.Password != "" && u.Name != "" && u.Create != 0 && u.LastLogin != 0
}

// ProtoToUser : turn protobuf to user
func ProtoToUser(user *protocol.UserInfo) *model.User {
	sex := 0
	if user.Sex {
		sex = 1
	}

	p := ""
	if len(user.Password)!=0{
		p = utils.MD5(user.Password)
	}

	return &model.User{
		Id:        user.Id,
		Role:      int(user.Role),
		Name:      user.Name,
		Sex:       sex,
		Phone:     user.Phone,
		Email:     user.Email,
		School:    user.School,
		LastLogin: user.LastLogin,
		Create:    user.Create,
		Account:   user.Account,
		Password:  p,
	}
}
