package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// BackStageManage : the backstage administration  module
type BackStageManage interface {
	// user manage
	GetUsers(r *http.Request) proto.Message
	AddUsers(r *http.Request) proto.Message
	UpdateUsers(r *http.Request) proto.Message
	DelUsers(r *http.Request) proto.Message

	// problems manage
	GetProblems(r *http.Request) proto.Message
	GetProblemByID(r *http.Request) proto.Message
	AddProblem(r *http.Request) proto.Message
	EditProblem(r *http.Request) proto.Message

	// class manage
	GetClasses(r *http.Request) proto.Message
	GetClassByID(r *http.Request) proto.Message
}
