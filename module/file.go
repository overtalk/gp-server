package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// FileItem : file item
type FileItem struct {
	ID string
	TS int64
}

// File : file module
type File interface {
	Daemon()
	Upload(r *http.Request) proto.Message
}
