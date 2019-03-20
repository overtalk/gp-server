package db

import (
	"errors"
)

var (
	// ErrNoMysqlConf : missing mysql configuration
	ErrNoMysqlConf = errors.New("lack of mysql config")
	// ErrNoRowsFound describes now rows found error
	ErrNoRowsFound = errors.New("no rows found")
	// ErrNoRowsAffected describes now rows affected error
	ErrNoRowsAffected = errors.New("no rows affected")
	// ErrMissingDefaultValue : missing default value
	ErrMissingDefaultValue = errors.New("some fields has no default value")
)
