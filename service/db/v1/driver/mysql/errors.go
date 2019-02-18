package driver

import (
	"errors"
)

var (
	// ErrNoRowsAffected describes now rows affected error
	ErrNoRowsAffected = errors.New("no rows affected")
	// ErrMissDocument describes no document set
	ErrMissDocument = errors.New("failed to set data : miss document")
	// ErrMissColumn describes no column set for db query
	ErrMissColumn = errors.New("failed to get data : miss column")
	// ErrMissData describes the error of lack data to exec
	ErrMissData = errors.New("failed to set data : miss data")
	// ErrNoRowsFound describes now rows found error
	ErrNoRowsFound = errors.New("no rows found")
)
