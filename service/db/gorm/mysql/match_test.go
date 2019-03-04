package db_test

import (
	"testing"
)

func TestMysqlDriver_GetMatchByInvitationCode(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	match, err := mysqlDriver.GetMatchByID(1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", match)
}
