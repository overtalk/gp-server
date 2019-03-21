package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Handler : handler func format
type Handler func(r *http.Request) proto.Message

// Router : method and handler
type Router struct {
	Method  string
	Handler Handler
}

// Gate : gateway module
type Gate interface {
	RegisterRoute(router, method string, handler Handler)
}
