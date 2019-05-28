package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Match : match module
type Match interface {
	NewMatch(r *http.Request) proto.Message
	GetMatches(r *http.Request) proto.Message
	GetMatchByID(r *http.Request) proto.Message
	GetPaperByID(r *http.Request) proto.Message
	EditMatch(r *http.Request) proto.Message
}
