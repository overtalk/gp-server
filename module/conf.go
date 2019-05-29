package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Conf : system conf module
type Conf interface {
	// Register(r *http.Request) proto.Message
	// Login(r *http.Request) proto.Message
	// Logout(r *http.Request) proto.Message

	GetConfig(r *http.Request) proto.Message
	GetUserRole(r *http.Request) proto.Message
	GetAllLanguage(r *http.Request) proto.Message
	GetJudgeResult(r *http.Request) proto.Message
	GetAlgorithm(r *http.Request) proto.Message

	// tag
	GetTags(r *http.Request) proto.Message
	AddTag(r *http.Request) proto.Message
	UpdateTag(r *http.Request) proto.Message
}
