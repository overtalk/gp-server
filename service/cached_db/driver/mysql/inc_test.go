package driver

import (
	"database/sql"
	"sync"
	"testing"
)

func TestInc(t *testing.T) {
	type args struct {
		p        *Pool
		document string
		columns  []string
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
				[]string{"cup", "season"},
				map[string]interface{}{
					"player_id": 2,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Inc(tt.args.p, tt.args.document, tt.args.columns, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("Inc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
