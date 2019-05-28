package module

import (
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

const (
	MANAGER = 3
)

// DB : database module
type DB interface {
	// conf
	GetAllDifficulty() ([]*model.Difficulty, error)
	GetAllLanguage() ([]*model.Language, error)
	GetAllTag() ([]*model.Tag, error)
	GetAllRole() ([]*model.Role, error)
	GetAlgorithm() ([]*model.Algorithm, error)

	// auth
	CreatePlayer(user *model.User) error
	CheckPlayer(username, password string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)

	// user manage
	GetUsersNum(role int64) (int64, error)
	GetUsers(pageNum, pageIndex int64) ([]*model.User, error)
	GetUsersByRole(pageNum, pageIndex, role int64) ([]*model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error // only id and changed filed is required
	DelUser(userID int64) error
	GetSubmitRecord(userID, problemID, pageNum, pageIndex int64) ([]*model.UserProblem, int64, error)

	// problem manage
	GetProblemsNum(tag int) (int64, error)
	GetProblemsByTagID(pageNum, pageIndex int64, tag int, keyword string) ([]*transform.IntactProblem, error)
	AddProblem(problem *transform.IntactProblem) error
	GetProblemByID(id int64) (*transform.IntactProblem, error)
	UpdateProblem(problem *transform.IntactProblem) error
	GetAllProblems() ([]*model.Problem, error) // 这个是用于智能组卷的接口，只需要获取部分信息（id，tags,通过率即可）

	// class manage
	GetClassNum() (int64, error)                                                           // 获得班级数量
	GetClasses(pageNum, pageIndex int64, keyword string) ([]*transform.IntactClass, error) // 获取所有班级
	AddClass(intactClass *transform.IntactClass) error                                     // 新增班级
	GetClassByID(id int64) (*transform.IntactClass, error)                                 // 获取班级具体信息
	UpdateClass(intactClass *transform.IntactClass) error                                  // 修改班级信息
	MemberManage(manageType, classID, memberID int64) error                                // 班级成员管理
	GetMembers(classID, pageNum, pageIndex int64) ([]*transform.UserClass, int64, error)   // 获取所有班级成员
	EnterClass(userID, classID int64) error                                                // 加入班级
	QuitClass(userID, classID int64)                                                       // 退出班级
	ApplyEnterRequest(userID, classID int64, isApply bool) error                           // 教师处理学生进入班级请求

	// rank
	GetRank(num int) ([]*RankItem, error)                        // 从数据库中读取排名
	GetNameAndSubmitNumByID(userID int64) (string, int64, error) // 查询用户姓名以及提交总数

	// announcement
	GetGlobalAnnouncementsNum() (int64, error)                                                  // 获得全局公告数量
	GetGlobalAnnouncements(pageNum, pageIndex int64) ([]*transform.AnnouncementWithName, error) // 获得全局公告
	GetAnnouncementsByClassID(classID int64) ([]*transform.AnnouncementWithName, error)         // 获得班级公告
	GetAnnouncementDetail(id int64) (*transform.AnnouncementWithName, error)                    // 获得公告具体信息
	AddAnnouncement(Announcement *model.Announcement) error                                     // 新增全局公告
	EditAnnouncement(Announcement *model.Announcement) error                                    // 修改全局公告
	DelAnnouncement(id int64) error                                                             // 删除全局公告

	// match
	GetMatchesNum() (int64, error)                               // 比赛数量
	AddMatch(match *model.Match) error                           // 构造新比赛
	GetMatches(pageNum, pageIndex int64) ([]*model.Match, error) // 获得所有比赛信息
	GetMatchByID(id int64) (*model.Match, error)                 // 获得比赛
	GetPaperByID(id int64) (*transform.Paper, error)             // 获取比赛的试卷信息
	EditMatch(match *model.Match) error                          // 修改比赛的信息

	// analysis
	GetDifficultyAnalysis(userID, startTime, endTime int64) (map[int64]float64, map[int64]int64, error)
	GetTagsAnalysis(userID, startTime, endTime int64, tags []int64) (map[int64]float64, map[int64]int64, error)

	// submit
	AddSubmitRecord(submit *model.UserProblem) error
}
