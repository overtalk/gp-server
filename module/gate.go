package module

import (
	"github.com/gin-gonic/gin"
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
