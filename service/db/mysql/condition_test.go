package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/db/mysql"
)

func TestTurnMysqlWhere(t *testing.T) {
	var conditions []db.Condition
	conditions = append(conditions, db.Condition{
		Filed:    "id",
		Operator: "<",
		Value:    1,
	})
	conditions = append(conditions, db.Condition{
		Filed:    "name",
		Operator: "=",
		Value:    "sf",
	})

	t.Log(db.TurnMysqlWhere(conditions))
}

func TestTurnMysqlWhere_nil(t *testing.T) {
	t.Log(db.TurnMysqlWhere(nil))
}
