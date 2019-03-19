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

// Router : method and handler
type Router struct {
	Method  string
	Handler Handler
}

// Gate : gateway module
type Gate interface {
	RegisterRoute(router, method string, handler Handler)
}
