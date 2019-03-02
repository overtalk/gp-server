package db

import (
	"errors"
)

var (
	// ErrNoMysqlConf : missing mysql configuration
	ErrNoMysqlConf = errors.New("lack of mysql config")
)
