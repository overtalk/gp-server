package source

import (
	"errors"
)

var (
	// ErrGetConfFail descirbes the error of get config fail, but get error code details
	ErrGetConfFail = errors.New("failed to get config from github")
)
