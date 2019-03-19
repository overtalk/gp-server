package dataError

import "errors"

var (
	// ErrNoRowsFound describes now rows found error
	ErrNoRowsFound = errors.New("no rows found")
	// ErrInvalidInput describes nil pointer for update error
	ErrInvalidInput = errors.New("invalid input parameters for update database")

	ErrInvalidType = errors.New("data type error")
)
