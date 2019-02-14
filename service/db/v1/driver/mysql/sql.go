package driver

import (
	"strings"

	"github.com/QHasaki/Server/model/v1"
)

// GetQuerySQL returns query sql
// the input arg `document` : cannot be nil
// the input arg `column` : at least include one string, "*" means query all
// the input arg `where` : can be nil
// check above before use func `Get` & `GetOne`
func GetQuerySQL(document string, column []string, where model.Data) (string, []interface{}) {
	var (
		columns string
		wheres  []string
		values  []interface{}
	)
	sql := "SELECT "

	columns = strings.Join(column, ",")

	for k, v := range where {
		switch v.(type) {
		case model.Where:
			wheres = append(wheres, k+" "+v.(model.Where).Operator+" ?")
			values = append(values, v.(model.Where).Value)
		default:
			wheres = append(wheres, k+" = ?")
			values = append(values, v)
		}
	}

	sql += columns + " FROM `" + document + "`"
	if len(wheres) != 0 {
		sql += " WHERE " + strings.Join(wheres, " AND ")
	}

	return sql, values
}

// GetExecSQL returns exec sql
// the input arg `document` : cannot be nil
// the input arg `data` : cannot be nil
// the input arg `where` : can be nil
// check above before use func `Get` & `GetOne`
func GetExecSQL(document string, data model.Data, where model.Data) (string, []interface{}) {
	// where == nil : create a new record
	if where == nil {
		var (
			column []string
			needed []string
			values []interface{}
		)

		sql := "INSERT INTO `" + document + "` "

		for k, v := range data {
			column = append(column, k)
			needed = append(needed, "?")
			values = append(values, v)
		}

		return sql + "(" + strings.Join(column, ",") + ") VALUES (" + strings.Join(needed, ", ") + ")", values
	}

	// where != nil : update old record
	var (
		datas  []string
		wheres []string
		values []interface{}
	)

	sql := "UPDATE `" + document + "` SET "

	for k, v := range data {
		datas = append(datas, k+" = ?")
		values = append(values, v)
	}

	for k, v := range where {
		switch v.(type) {
		case model.Where:
			wheres = append(wheres, k+" "+v.(model.Where).Operator+" ?")
			values = append(values, v.(model.Where).Value)
		default:
			wheres = append(wheres, k+" = ?")
			values = append(values, v)
		}
	}

	sql += strings.Join(datas, ", ")

	if len(wheres) != 0 {
		sql += " WHERE " + strings.Join(wheres, " AND ")
	}

	return sql, values

}
