package data

import (
	"strconv"
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/service/cached_db/cache/memory"
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
)

func TestNewConnect(t *testing.T) {
	dbinfo := DataInfo{
		Size:     10,
		Addr:     "172.26.32.12",
		Username: "sausage",
		Password: "sausage_shooter",
		DBName:   "sausage_shooter",
	}
	_, err := NewConnect(dbinfo, CacheInfo{})
	if err != nil {
		t.Fatalf("Test Data NewConnect faild")
	}
}

func TestDataDB_Set(t *testing.T) {
	dbinfo := DataInfo{
		Size:     10,
		Addr:     "172.26.32.12",
		Username: "sausage",
		Password: "sausage_shooter",
		DBName:   "sausage_shooter",
	}
	db, err := NewConnect(dbinfo, CacheInfo{})
	if err != nil {
		t.Fatalf("data NewConnect faild")
	}

	data := make(Data)
	data["id"] = "test_data"
	data["nickname"] = "test_name_" + strconv.Itoa(time.Now().Second())
	data["url"] = "http://127.0.0.1/demo.png"

	where := make(Data)
	where["id"] = "test_data"

	err = db.Set("player", data, where)
	if err != nil {
		t.Fatalf("Test Data Failed: %v", err)
	}
}

func TestDataDB_Get(t *testing.T) {
	dbinfo := DataInfo{
		Size:     10,
		Addr:     "172.26.32.12",
		Username: "sausage",
		Password: "sausage_shooter",
		DBName:   "sausage_shooter",
	}
	db, err := NewConnect(dbinfo, CacheInfo{})
	if err != nil {
		t.Fatalf("data NewConnect faild")
	}
	where := make(Data)
	where["id"] = "test_data"

	data, err := db.Get("player", []string{"*"}, where)
	if err != nil {
		t.Fatalf("Test Data Failed: %v", err)
	}

	t.Logf("Test Data Get : %v", data)
}

func TestDB_NoCache(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *originDB
	}{
		{
			"function test",
			fields{
				driver.DBInfo{
					Size:     10,
					Addr:     "localhost",
					Username: "root",
					Password: "12341234",
					DBName:   "sausage_shooter",
				},
				cache.CacheInfo{},
			},
			&originDB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			_ = p.NoCache()
		})
	}
}
