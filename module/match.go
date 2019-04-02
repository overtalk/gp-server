package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Match : match module
type Match interface {
	NewMatch(r *http.Request) proto.Message
}
