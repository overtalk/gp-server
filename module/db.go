package module

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// DB : database module
type DB interface {
	// auth
	CheckPlayer(username, password string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)

	// user manage
	GetUsersNum() (int64, error)
	GetUsers(pageNum, pageIndex int64) ([]*model.User, error)
	GetUsersByRole(pageNum, pageIndex, role int64) ([]*model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error // only id and changed filed is required
	DelUser(userID int64) error

	// problem manage
	GetProblemsNum() (int64, error)
	GetProblems(pageNum, pageIndex int64) ([]*transform.IntactProblem, error)
	GetProblemsByTagID(pageNum, pageIndex int64, tag int) ([]*transform.IntactProblem, error)
	AddProblem(problem *transform.IntactProblem) error
	GetProblemByID(id int64) (*transform.IntactProblem, error)
	UpdateProblem(problem *transform.IntactProblem) error

	// class manage
	GetClassNum() (int64, error)                                           // 获得班级数量
	GetClasses(pageNum, pageIndex int64) ([]*transform.IntactClass, error) // 获取所有班级
	AddClass(intactClass *transform.IntactClass) error                     // 新增班级
	GetClassByID(id int64) (*transform.IntactClass, error)                 // 获取班级具体信息
	UpdateClass(intactClass *transform.IntactClass) error                  // 修改班级信息
	MemberManage(manageType, classID, memberID int64) error                // 班级成员管理

	// rank
	GetRank(num int) ([]*RankItem, error) // 从数据库中读取排名
}
