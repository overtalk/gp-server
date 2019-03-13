package module

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// DB : database module
type DB interface {
	// auth
	CheckPlayer(username, password string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)

	// user manage
	GetUsersByRole(role int64) ([]*model.User, error) // role < 0 : get all user

	GetMatchByID(id int) (*model.Match, error)
}
