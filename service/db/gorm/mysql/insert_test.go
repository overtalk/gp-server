package db_test

import (
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/model/gorm"
)

func TestMysqlDriver_Insert(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	test := &model.Test{
		Nickname:    "qinhan",
		CreatedAt:   time.Now(),
		Achievement: []byte("aaaa"),
		Level:       3,
	}

	mysqlDriver.Insert(test)
}
