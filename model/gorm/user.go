package model

import (
	"github.com/qinhan-shu/gp-server/protocol"
)

// User : teble `user`
type User struct {
	ID        int64  `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Account   string `gorm:"type : varchar(50); not null; unique"`
	Password  string `gorm:"type : varchar(100); not null"`
	Role      int    `gorm:"type : tinyint(4); not null; default : 0"`
	Name      string `gorm:"type : varchar(50); not null"`
	Sex       bool   `gorm:"type : boolean; not null; default : false"`
	Phone     string `gorm:"type : varchar(20)"`
	Email     string `gorm:"type : varchar(50)"`
	School    string `gorm:"type : varchar(50)"`
	Create    int64  `gorm:"type : int(64); not null"`
	LastLogin int64  `gorm:"type : int(64); not null"`
}

// TurnProto : turn user to protobuf
// username & password is not allowed to send to client
func (u *User) TurnProto() *protocol.UserInfo {
	return &protocol.UserInfo{
		Id:        u.ID,
		Role:      protocol.Role(u.Role),
		Name:      u.Name,
		Sex:       u.Sex,
		Phone:     u.Phone,
		Email:     u.Email,
		School:    u.School,
		Create:    u.Create,
		LastLogin: u.LastLogin,
	}
}

// IsInited : check the default value of each fields
func (u *User) IsInited() bool {
	return u.Account != "" && u.Password != "" && u.Name != "" && u.Create != 0 && u.LastLogin != 0
}

// TurnUser : turn protobuf to user
func TurnUser(user *protocol.UserInfo) *User {
	return &User{
		ID:        user.Id,
		Role:      int(user.Role),
		Name:      user.Name,
		Sex:       user.Sex,
		Phone:     user.Phone,
		Email:     user.Email,
		School:    user.School,
		LastLogin: user.LastLogin,
		Create:    user.Create,
		Account:   user.Account,
		Password:  user.Password,
	}
}
