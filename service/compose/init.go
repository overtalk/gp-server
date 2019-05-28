package compose

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Compose : implementation of compose module
type Compose struct {
	db module.DB
}

// NewCompose : constructor for paper match
func NewCompose(dataStorage *module.DataStorage) module.Compose {
	return &Compose{
		db: dataStorage.DB,
	}
}
