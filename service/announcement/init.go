package announcement

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Announcement : implementation of auth announcement
type Announcement struct {
	cache module.Cache
	db    module.DB
}

// NewAnnouncement : constructor for module Announcement
func NewAnnouncement(dataStorage *module.DataStorage) module.Announcement {
	return &Announcement{
		cache: dataStorage.Cache,
		db:    dataStorage.DB,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewAnnouncement(dataStorage)
	gate.RegisterRoute("/getAnnouncements", "POST", module.GetAnnouncements)
	gate.RegisterRoute("/announcementDetail", "POST", module.GetDetail)
	gate.RegisterRoute("/addAnnouncement", "POST", module.AddAnnouncement)
	gate.RegisterRoute("/editAnnouncement", "POST", module.EditAnnouncement)
	gate.RegisterRoute("/delAnnouncement", "POST", module.DelAnnouncement)
}
