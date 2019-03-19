package cache_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/qinhan-shu/gp-server/service/cached_db/cache/memory"
)

func TestCache(t *testing.T) {
	cache, err := cache.Connect(cache.CacheInfo{})
	if err != nil {
		t.Fatalf("failed to connect a cache pool : %v", err)
	}
	testData := make(map[string]interface{})
	testData["string"] = "aa"
	testData["int"] = 10
	testData["time"] = time.Now()
	testData["bool"] = true
	testData["float"] = 10.10

	err = cache.SetAll("testData", testData)
	if err != nil {
		t.Fatalf("failed to set test data in cache : %v", err)
	}

	gotData, err := cache.GetAll("testData")
	if err != nil {
		t.Fatalf("failed to got test data in cache : %v", err)
	}

	if reflect.DeepEqual(testData, gotData) {
		t.Fatalf("failed equal test data")
	}

	err = cache.Delete("testData")
	if err != nil {
		t.Fatalf("failed to delete a test data in cache : %v", err)
	}

	_, err = cache.GetAll("testData")
	if err == nil {
		t.Fatalf("got a deleted test data in cache")
	}
}
