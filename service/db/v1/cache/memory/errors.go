package cache

import "errors"

var (
	// ErrNoRowsFound describes now rows found error
	ErrNoRowsFound = errors.New("no rows found")
	// ErrInvalidType describes invalid variable type
	ErrInvalidType = errors.New("data type error")
)
