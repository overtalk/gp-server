package data

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/cached_db/cache/memory"
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
)

func TestDB_CacheRefresh(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	type args struct {
		document string
		where    Data
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
				driver.DBInfo{
					Size:     10,
					Addr:     "localhost",
					Username: "root",
					Password: "12341234",
					DBName:   "sausage_shooter",
				},
				cache.CacheInfo{},
			},
			args{
				"player_cup",
				map[string]interface{}{
					"id":      1,
					"open_id": "syj",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			if err := p.CacheRefresh(tt.args.document, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("DB.CacheRefresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_CacheDaemon(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	tests := []struct {
		name   string
		fields fields
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			p.CacheDaemon()
		})
	}
}
