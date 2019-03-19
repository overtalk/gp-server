package db_test

import (
	"testing"
)

func TestMysqlDriver_Get(t *testing.T) {
	mysql, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	result, err := mysql.Get("user", []string{"*"}, "")
	if err != nil {
		t.Error(err)
		return
	}

	r := result.(map[string]interface{})
	for k, v := range r {
		t.Logf("%s = %v", k, v)
	}
}
