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
	GetUsers(pageNum, pageIndex int64) ([]*model.User, error)
	GetUsersByRole(pageNum, pageIndex, role int64) ([]*model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error // only id and changed filed is required
	DelUser(userID int64) error

	// problem manage
	GetProblems(pageNum, pageIndex int64) ([]*model_utils.IntactProblem, error)
	GetProblemsByTagID(pageNum, pageIndex int64, tag int) ([]*model_utils.IntactProblem, error)
	AddProblem(problem *model_utils.IntactProblem) error
	GetProblemByID(id int64) (*model_utils.IntactProblem, error)
	UpdateProblem(problem *model_utils.IntactProblem) error

	// tags

}
