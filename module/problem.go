package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Problem : problem  module
type Problem interface {
	GetProblems(r *http.Request) proto.Message
	GetProblemByID(r *http.Request) proto.Message
	AddProblem(r *http.Request) proto.Message
	EditProblem(r *http.Request) proto.Message
}
