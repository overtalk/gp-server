package db

import (
	"errors"
)

var (
	// ErrNoMysqlConf : missing mysql configuration
	ErrNoMysqlConf = errors.New("lack of mysql config")
	// ErrMissingDefaultValue : missing default value
	ErrMissingDefaultValue = errors.New("some fields has no default value")
)
