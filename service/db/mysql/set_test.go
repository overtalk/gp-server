package db_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/service/db/mysql"
)

func TestMysqlDriver_Update(t *testing.T) {
	mysql, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	document := "user"

	data := make(map[string]interface{})
	data["name"] = "TestMysqlDriver_Set"

	if err := mysql.Set(document, data, db.Condition{
		Filed:    "id",
		Operator: "=",
		Value:    1,
	}); err != nil {
		t.Error(err)
		return
	}
}

func TestMysqlDriver_Insert(t *testing.T) {
	mysql, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	document := "user"

	data := make(map[string]interface{})
	data["account"] = "jac"
	data["password"] = "jack"
	data["name"] = "jack"
	data["create"] = time.Now().Unix()
	data["last_login"] = time.Now().Unix()

	if err := mysql.Set(document, data); err != nil {
		t.Error(err)
		return
	}
}
