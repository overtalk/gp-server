package auth

import (
	"github.com/qinhan-shu/gp-server/module"
)

// Auth : implementation of auth module
type Auth struct {
	db    module.DB
	cache module.Cache
}

// NewAuth : constructor for module Auth
func NewAuth(dataStorage module.DataStorage) *Auth {
	return &Auth{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage module.DataStorage) {
	// authModule := NewAuth(dataStorage)
	// gate.RegisterRoute("/example", module.Router{
	// 	Method:  "POST",
	// 	Handler: authModule,
	// })
}
