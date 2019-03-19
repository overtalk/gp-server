package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/db/mysql"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

func TestMysqlDriver_Get(t *testing.T) {
	mysql, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	result, err := mysql.Get("user", []string{"*"})
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("id[%d] --- %v", parse.Int(result["id"]), result)
}

func TestMysqlDriver_Gets(t *testing.T) {
	mysql, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	results, err := mysql.Gets("user", []string{"*"}, db.Condition{
		Filed:    "id",
		Operator: ">",
		Value:    10,
	})
	if err != nil {
		t.Error(err)
		return
	}

	for _, result := range results {
		t.Logf("id[%d] --- %v", parse.Int(result["id"]), result)
	}
}
