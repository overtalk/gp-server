package db

import (
	"errors"
)

var (
	// ErrNoMysqlConf : missing mysql configuration
	ErrNoMysqlConf = errors.New("lack of mysql config")
	// ErrNoRowsAffected describes now rows affected error
	ErrNoRowsAffected = errors.New("no rows affected")
	// ErrInvaildGetArgs describes error of Get input args error
	ErrInvaildGetArgs = errors.New("get input args error")
)
