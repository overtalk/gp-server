package module

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// DB : database module
type DB interface {
	GetUserByID(id int) (*model.User, error)
	GetMatchByID(id int) (*model.Match, error)
}
