package module

import (
	"github.com/gin-gonic/gin"
)

// Auth : User identity authentication module
type Auth interface {
	Login(c *gin.Context) interface{}
	Logout(c *gin.Context) interface{}
}
