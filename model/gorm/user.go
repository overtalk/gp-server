package model

import (
	"github.com/qinhan-shu/gp-server/protocol"
)

// User : teble `user`
type User struct {
	ID            int64  `gorm:"type : int(20); primary_key; not null; auto_increment"`
	Username      string `gorm:"type : varchar(50); not null; unique"`
	Password      string `gorm:"type : varchar(100); not null"`
	OperationAuth int    `gorm:"type : tinyint(4); not null; default : 0"`
	Role          int    `gorm:"type : tinyint(4); not null; default : 0"`
	Name          string `gorm:"type : varchar(50); not null"`
	Sex           bool   `gorm:"type : boolean; not null"`
	Email         string `gorm:"type : varchar(50); not null"`
	Academy       string `gorm:"type : varchar(50); not null"`
	Major         string `gorm:"type : varchar(50); not null"`
	Create        int64  `gorm:"type : int(64); not null"`
	LastLogin     int64  `gorm:"type : int(64); not null"`
}

// TurnProto : turn user to protobuf
func (u *User) TurnProto() *protocol.UserInfo {
	return &protocol.UserInfo{
		Id:        u.ID,
		Name:      u.Name,
		Sex:       u.Sex,
		Role:      protocol.Role(u.Role),
		Academy:   u.Academy,
		Major:     u.Major,
		Create:    u.Create,
		LastLogin: u.LastLogin,
	}
}
