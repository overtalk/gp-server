package module

import (
	"github.com/gin-gonic/gin"
)

// input args key in the map of Handler's args
var (
	Token   = "1"
	Request = "2"
)

// Handler : handler func format
type Handler func(c *gin.Context) interface{}

// Gate : gateway module
type Gate interface {
	RegisterRoute(router string, handler Handler)
}
