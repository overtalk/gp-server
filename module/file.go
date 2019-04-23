package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// File : file module
type File interface {
	Upload(r *http.Request) proto.Message
}
