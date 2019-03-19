package data

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/cached_db/cache/memory"
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

func TestDB_Set(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	type args struct {
		document string
		data     Data
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
					"season": 2,
				},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			if err := p.Set(tt.args.document, tt.args.data, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("DB.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Get(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	type args struct {
		document string
		column   []string
		where    Data
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Data
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
				[]string{"cup"},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			map[string]interface{}{
				"cup":       2446,
				"season":    2,
				"player_id": 1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			got, err := p.Get(tt.args.document, tt.args.column, tt.args.where)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k, v := range got {
				if parse.String(v) != parse.String(tt.want[k]) {
					t.Errorf("DB.Get() v %v, wantValue %v", v, tt.want[k])
					break
				}
			}
		})
	}
}

func TestDB_Inc(t *testing.T) {
	type fields struct {
		dataInfo  DataInfo
		cacheInfo CacheInfo
	}
	type args struct {
		document string
		column   []string
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
				[]string{"cup", "season"},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := NewConnect(tt.fields.dataInfo, tt.fields.cacheInfo)
			if err := p.Inc(tt.args.document, tt.args.column, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("DB.Inc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
