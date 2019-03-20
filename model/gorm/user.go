package model

import (
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

var (
	UserID        = "id"
	UserAccount   = "account"
	UserPassword  = "password"
	UserRole      = "role"
	UserName      = "name"
	UserSex       = "sex"
	UserPhone     = "phone"
	UserEmail     = "email"
	UserSchool    = "school"
	UserCreate    = "create"
	UserLastLogin = "last_login"
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
		Password:  utils.MD5(user.Password),
	}
}

// TurnMapToProtoUser : turn map to protobuf user
func TurnMapToProtoUser(data map[string]interface{}) (*User, error) {
	// id
	id, err := parse.IntWithError(data[UserID])
	if err != nil {
		return nil, err
	}
	// account
	account, err := parse.StringWithError(data[UserAccount])
	if err != nil {
		return nil, err
	}
	// role
	role, err := parse.IntWithError(data[UserRole])
	if err != nil {
		return nil, err
	}
	// name
	name, err := parse.StringWithError(data[UserName])
	if err != nil {
		return nil, err
	}
	// sex
	sex, err := parse.IntWithError(data[UserSex])
	if err != nil {
		return nil, err
	}
	// phone
	phone, err := parse.StringWithError(data[UserPhone])
	if err != nil {
		return nil, err
	}
	// email
	email, err := parse.StringWithError(data[UserEmail])
	if err != nil {
		return nil, err
	}
	// school
	school, err := parse.StringWithError(data[UserSchool])
	if err != nil {
		return nil, err
	}
	// create
	create, err := parse.IntWithError(data[UserCreate])
	if err != nil {
		return nil, err
	}
	// last_login
	lastLogin, err := parse.IntWithError(data[UserLastLogin])
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        id,
		Account:   account,
		Role:      int(role),
		Name:      name,
		Sex:       sex == 1,
		Phone:     phone,
		Email:     email,
		School:    school,
		Create:    create,
		LastLogin: lastLogin,
	}, nil
}
