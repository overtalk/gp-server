package module

import (
	model_utils "github.com/qinhan-shu/gp-server/model"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// DB : database module
type DB interface {
	// auth
	CheckPlayer(username, password string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)

	// user manage
	GetUsersByRole(role int64) ([]*model.User, error) // role < 0 : get all user
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error // only id and changed filed is required
	DelUser(userID int64) error

	// problem manage
	GetProblems() ([]*model_utils.IntactProblem, error)
	GetProblemsByTagID(tag int) ([]*model_utils.IntactProblem, error)
	AddProblem(problem *model_utils.IntactProblem) error
	GetProblemByID(id int64) (*model_utils.IntactProblem, error)
	UpdateProblem(problem *model_utils.IntactProblem) error

	// tags

}
