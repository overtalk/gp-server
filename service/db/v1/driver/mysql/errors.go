package driver

import (
	"errors"
)

var (
	// ErrNoRowsAffected describes now rows affected error
	ErrNoRowsAffected = errors.New("no rows affected")
)
