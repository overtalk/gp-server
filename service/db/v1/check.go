package db

import (
	"github.com/QHasaki/Server/model/v1"
)

// set illegal condition :
// document == ""
// len(data) < 1
func checkSetCondition(document string, data model.Data, where model.Data) error {
	if document == "" {
		return ErrMissDocument
	}
	if len(data) < 1 {
		return ErrMissData
	}
	return nil
}

// get illegal condition :
// document == ""
// len(where) < 1
// len(column) < 1
func checkGetCondition(document string, column []string, where model.Data) error {
	if document == "" {
		return ErrMissDocument
	}
	if len(where) < 1 {
		return ErrMissWhere
	}
	if len(column) < 1 {
		return ErrMissColumn
	}
	return nil
}
