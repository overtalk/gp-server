package module

import (
	"github.com/gin-gonic/gin"
)

// BackStageManage : the backstage administration  module
type BackStageManage interface {
	// user manage
	GetUsers(c *gin.Context) interface{}
	AddUsers(c *gin.Context) interface{}
	UpdateUsers(c *gin.Context) interface{}
	DelUsers(c *gin.Context) interface{}

	// problems manage
	GetProblems(c *gin.Context) interface{}
	GetProblemByID(c *gin.Context) interface{}
	AddProblem(c *gin.Context) interface{}
	EditProblem(c *gin.Context) interface{}
}
