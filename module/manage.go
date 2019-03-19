package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// BackStageManage : the backstage administration  module
type BackStageManage interface {
	// user manage
	GetUsers(r *http.Request) (int, proto.Message)
	AddUsers(r *http.Request) (int, proto.Message)
	UpdateUsers(r *http.Request) (int, proto.Message)
	DelUsers(r *http.Request) (int, proto.Message)

	// problems manage
	GetProblems(r *http.Request) (int, proto.Message)
	GetProblemByID(r *http.Request) (int, proto.Message)
	AddProblem(r *http.Request) (int, proto.Message)
	EditProblem(r *http.Request) (int, proto.Message)
}
