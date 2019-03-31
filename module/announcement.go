package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Announcement : global announcement module
type Announcement interface {
	GetAnnouncements(r *http.Request) proto.Message
	GetDetail(r *http.Request) proto.Message
}
