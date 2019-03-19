package db

import (
	"strings"
)

// Condition : db set/get condition
type Condition struct {
	Filed    string
	Operator string
	Value    interface{}
}

// TurnMysqlWhere : turn to sql
func TurnMysqlWhere(conditions []Condition) (string, []interface{}) {
	var where []string
	var args []interface{}
	for _, condition := range conditions {
		where = append(where, condition.Filed+" "+condition.Operator+" ?")
		args = append(args, condition.Value)
	}
	return strings.Join(where, " AND "), args
}
