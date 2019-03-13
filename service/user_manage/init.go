package user_manage

import (
	"github.com/qinhan-shu/gp-server/module"
)

// UserManager : implementation of UserManage module
type UserManager struct {
	db    module.DB
	cache module.Cache
}

// NewUserManager : constructor for module NewUserManager
func NewUserManager(dataStorage *module.DataStorage) *UserManager {
	return &UserManager{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	userManagerModule := NewUserManager(dataStorage)
	gate.RegisterRoute("/getUsers", module.Router{
		Method:  "POST",
		Handler: userManagerModule.GetUsers,
	})
}
