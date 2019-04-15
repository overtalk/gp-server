package transform

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// UserClass : user in class
type UserClass struct {
	model.UserClass `xorm:"extends"`
	Name            string
}

func (UserClass) TableName() string {
	return "user_class"
}

// UserClassToProto : turn announcement to protobuf
func UserClassToProto(u *UserClass) *protocol.ClassMember {
	return &protocol.ClassMember{
		UserId:          u.UserId,
		IsAdministrator: u.IsAdministrator == 1,
		IsChecked:       u.IsChecked == 1,
		Name:            u.Name,
	}
}
