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
func NewAuth(dataStorage *module.DataStorage) module.Auth {
	return &Auth{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewAuth(dataStorage)
	gate.RegisterRoute("/login", "POST", module.Login)
	gate.RegisterRoute("/logout", "GET", module.Logout)
	gate.RegisterRoute("/conf", "GET", module.GetConfig)
	gate.RegisterRoute("/userRole", "GET", module.GetUserRole)
}
