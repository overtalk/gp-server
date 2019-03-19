package module

import (
	"github.com/gin-gonic/gin"
)

// BackStageManage : the backstage administration  module
type BackStageManage interface {
	// user manage
	GetUsers(c *gin.Context) (int, interface{})
	AddUsers(c *gin.Context) (int, interface{})
	UpdateUsers(c *gin.Context) (int, interface{})
	DelUsers(c *gin.Context) (int, interface{})

	// problems manage
	GetProblems(c *gin.Context) (int, interface{})
	GetProblemByID(c *gin.Context) (int, interface{})
	AddProblem(c *gin.Context) (int, interface{})
	EditProblem(c *gin.Context) (int, interface{})
}
