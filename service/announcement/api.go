package announcement

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (a *Announcement) GetAnnouncements(r *http.Request) proto.Message {
	return nil
}

func (a *Announcement) GetDetail(r *http.Request) proto.Message {
	return nil
}
