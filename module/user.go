package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// User : user  module
type User interface {
	GetUsers(r *http.Request) proto.Message
	AddUsers(r *http.Request) proto.Message
	UpdateUsers(r *http.Request) proto.Message
	DelUsers(r *http.Request) proto.Message
	GetSubmitRecord(r *http.Request) proto.Message
}
