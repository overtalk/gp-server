package user

import (
	"github.com/qinhan-shu/gp-server/module"
)

// User : implementation of backstage user module
type User struct {
	judgeFilePath string
	db            module.DB
	cache         module.Cache
}

// NewUser : constructor for module User
func NewUser(dataStorage *module.DataStorage) module.User {
	return &User{
		// judgeFilePath: dataStorage.JudgeFilePath,
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewUser(dataStorage)
	// user manage
	gate.RegisterRoute("/getUsers", "POST", module.GetUsers)
	gate.RegisterRoute("/addUsers", "POST", module.AddUsers)
	gate.RegisterRoute("/updateUsers", "POST", module.UpdateUsers)
	gate.RegisterRoute("/delUsers", "POST", module.DelUsers)
	gate.RegisterRoute("/submitRecord", "POST", module.GetSubmitRecord)
}
