package driver

import (
	"database/sql"
	"sync"
	"testing"
)

func TestSet(t *testing.T) {
	type args struct {
		p        *Pool
		document string
		data     Data
		where    Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"function test",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player_cup",
				map[string]interface{}{
					"season": 2,
				},
				map[string]interface{}{
					"player_id": "1",
				},
			},
			false,
		},
		{
			"where == nil",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player_cup",
				map[string]interface{}{
					"player_id": 2,
					"season":    1,
					"cup":       1,
				},
				nil,
			},
			false,
		},
		{
			"where.value.(type) == Where",
			args{
				&Pool{
					CorrectMysqlTestUrl,
					10,
					map[*sql.DB]bool{},
					new(sync.Mutex),
				},
				"player_cup",
				map[string]interface{}{
					"season": 2,
				},
				map[string]interface{}{
					"player_id": Where{
						Operator: "=",
						Value:    2,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Set(tt.args.p, tt.args.document, tt.args.data, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
