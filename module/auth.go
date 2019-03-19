package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Auth : User identity authentication module
type Auth interface {
	Login(r *http.Request) (int, proto.Message)
	Logout(r *http.Request) (int, proto.Message)
}
