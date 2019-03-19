package driver

import (
	"database/sql"
	"reflect"
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		p        *Pool
		document string
		column   []string
		where    Data
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			"function test1",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player",
				[]string{"open_id"},
				map[string]interface{}{
					"id": Where{
						Operator: "=",
						Value:    "1",
					},
				},
			},
			map[string]interface{}{
				"open_id": []byte("syj"),
			},
			false,
		},
		{
			"function test2",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player",
				[]string{"open_id"},
				map[string]interface{}{
					"id": "1",
				},
			},
			map[string]interface{}{
				"open_id": []byte("syj"),
			},
			false,
		},
		{
			"function test3",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player",
				[]string{"open_id"},
				map[string]interface{}{
					"id": Where{
						Operator: "=",
						Value:    "2",
					},
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.p, tt.args.document, tt.args.column, tt.args.where)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
