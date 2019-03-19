package cache

import (
	"sync"
	"testing"
	"time"
)

func TestCache_Delete(t *testing.T) {
	f := func(commonMap *map[interface{}]interface{}) sync.Map {
		syncMap := new(sync.Map)
		for k, v := range *commonMap {
			syncMap.Store(k, v)
		}
		return *syncMap
	}
	type fields struct {
		storage sync.Map
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"function test",
			fields{
				f(&map[interface{}]interface{}{
					"k": "v",
				}),
			},
			args{
				"k",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				storage: tt.fields.storage,
			}
			if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCache_Daemon(t *testing.T) {
	f := func(commonMap *map[interface{}]*CacheData) sync.Map {
		syncMap := new(sync.Map)
		for k, v := range *commonMap {
			syncMap.Store(k, v)
		}
		return *syncMap
	}
	type fields struct {
		storage sync.Map
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"测试有过期数据",
			fields{
				f(&map[interface{}]*CacheData{
					"1": {
						[]byte("test data"),
						time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
					},
				}),
			},
		},
		{
			"测试无过期数据",
			fields{
				f(&map[interface{}]*CacheData{
					"1": {
						[]byte("test data"),
						time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local),
					},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// daemon进程，不方便测试
			c := &Cache{
				storage: tt.fields.storage,
			}
			c.Daemon()
		})
	}
}
