package db_test

import (
	"testing"
	"time"
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

	if err := mysql.Set(document, data, "id = ?", 1); err != nil {
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
	data["account"] = "jack"
	data["password"] = "jack"
	data["name"] = "jack"
	data["create"] = time.Now().Unix()
	data["last_login"] = time.Now().Unix()

	if err := mysql.Set(document, data, ""); err != nil {
		t.Error(err)
		return
	}
}
