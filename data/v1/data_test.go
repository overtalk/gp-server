package data

import (
	"strconv"
	"testing"
	"time"
)

func TestNewConnect(t *testing.T) {
	dbinfo := DataInfo{
		Size:     10,
		Addr:     "172.26.32.12",
		Username: "sausage",
		Password: "sausage_shooter",
		DBName:   "sausage_shooter",
	}
	_, err := NewConnect(dbinfo)
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
	db, err := NewConnect(dbinfo)
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
	db, err := NewConnect(dbinfo)
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
