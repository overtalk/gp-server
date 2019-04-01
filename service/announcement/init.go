package announcement

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Announcement : implementation of auth announcement
type Announcement struct {
	db module.DB
}

// NewAnnouncement : constructor for module Announcement
func NewAnnouncement(dataStorage *module.DataStorage) module.Announcement {
	return &Announcement{
		db: dataStorage.DB,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewAnnouncement(dataStorage)
	gate.RegisterRoute("/getAnnouncements", "POST", module.GetAnnouncements)
	gate.RegisterRoute("/announcementDetail", "GET", module.GetDetail)
}
