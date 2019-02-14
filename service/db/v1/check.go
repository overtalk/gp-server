package db

import (
	"github.com/QHasaki/Server/model/v1"
)

func setCondition(document string, data model.Data, where model.Data) error {
	if document == "" {
		return ErrMissDocument
	}
	if len(data) < 1 {
		return ErrMissData
	}
	return nil
}

func getCondition(document string, column []string, where model.Data) error {
	if document == "" {
		return ErrMissDocument
	}
	if where == nil || len(where) < 1 {
		return ErrMissWhere
	}
	if len(column) < 1 {
		return ErrMissColumn
	}
	return nil
}
