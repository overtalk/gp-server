package db

import (
	"errors"
)

var (
	// ErrMissDocument describes no document set
	ErrMissDocument = errors.New("failed to set data : miss document")
	// ErrMissColumn describes no column set for db query
	ErrMissColumn = errors.New("failed to get data : miss column")
	// ErrMissData describes the error of lack data to exec
	ErrMissData = errors.New("failed to set data : miss data")
	// ErrMissWhere describes the error of lack of query condition
	ErrMissWhere = errors.New("failed to get data : miss where condition")
	// ErrMissPK describes the error of lack of pk in get cache key
	ErrMissPK = errors.New("failed to get cache key : miss pk")
)
