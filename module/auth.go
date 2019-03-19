package module

import (
	"github.com/gin-gonic/gin"
)

// Auth : User identity authentication module
type Auth interface {
	Login(c *gin.Context) (int, interface{})
	Logout(c *gin.Context) (int, interface{})
}
