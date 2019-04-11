package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Auth : User identity authentication module
type Auth interface {
	Login(r *http.Request) proto.Message
	Logout(r *http.Request) proto.Message
	GetConfig(r *http.Request) proto.Message
}
