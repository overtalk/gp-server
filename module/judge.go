package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// JudgeServerConf : config of judge server
type JudgeServerConf struct {
	Addr  string
	Token string
}

// Judge : judge module
type Judge interface {
	Judge(r *http.Request) proto.Message
}
