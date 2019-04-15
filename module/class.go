package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Class : class module
type Class interface {
	GetClasses(r *http.Request) proto.Message
	GetClassByID(r *http.Request) proto.Message
	AddClass(r *http.Request) proto.Message
	EditClass(r *http.Request) proto.Message
	MemberManage(r *http.Request) proto.Message
	GetMembers(r *http.Request) proto.Message
	EnterClass(r *http.Request) proto.Message
}
