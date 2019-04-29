package db_test

import (
	"testing"
	"time"
)

func TestMysqlDriver_GetDifficultyAnalysis(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var userID int64 = 1
	var startTime int64 = 0
	var endTime int64 = time.Now().Unix()
	line, pie, err := mysqlDriver.GetDifficultyAnalysis(userID, startTime, endTime)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("pass rate: %v", line)
	t.Logf("pass num: %v", pie)
}

func TestMysqlDriver_GetTagsAnalysis(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var userID int64 = 1
	var startTime int64 = 0
	var endTime int64 = time.Now().Unix()
	tags := []int64{1, 2, 3}
	line, pie, err := mysqlDriver.GetTagsAnalysis(userID, startTime, endTime, tags)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("pass rate: %v", line)
	t.Logf("pass num: %v", pie)
}
