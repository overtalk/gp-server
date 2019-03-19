package data

import (
	"errors"
)

func getCondition(document string, data Data, where Data) error {
	if document == "" {
		return errors.New("failed to set data : miss document")
	}
	if len(data) < 1 {
		return errors.New("failed to set data : miss data")
	}
	return nil
}

func setCondition(document string, column []string, where Data) error {
	if document == "" {
		return errors.New("failed to get data : miss document")
	}
	if where == nil || len(where) < 1 {
		return errors.New("failed to get data : miss where condition")
	}
	if len(column) < 1 {
		return errors.New("failed to get data : miss column")
	}
	return nil
}
