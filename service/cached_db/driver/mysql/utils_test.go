package driver

import (
	"database/sql"
	"sync"
	"testing"
)

func TestPool_Exec(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	type args struct {
		sql  string
		args []interface{}
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
				CorrectMysqlTestUrl,
				10,
				map[*sql.DB]bool{},
				new(sync.Mutex),
			},
			args{
				"update player_cup t set t.cup=t.cup+1 where player_id=?",
				[]interface{}{1},
			},
			false,
		},
		{
			"sql statement is nil",
			fields{
				CorrectMysqlTestUrl,
				10,
				map[*sql.DB]bool{},
				new(sync.Mutex),
			},
			args{
				"",
				[]interface{}{},
			},
			true,
		},
		{
			"affected row count = 0",
			fields{
				CorrectMysqlTestUrl,
				10,
				map[*sql.DB]bool{},
				new(sync.Mutex),
			},
			args{
				"update player_cup t set t.cup=t.cup+1 where player_id=?",
				[]interface{}{0},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				dsn:  tt.fields.dsn,
				size: tt.fields.size,
				dbs:  tt.fields.dbs,
				l:    tt.fields.l,
			}
			if err := p.Exec(tt.args.sql, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Pool.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPool_Query(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	type args struct {
		query string
		args  []interface{}
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
				CorrectMysqlTestUrl,
				10,
				map[*sql.DB]bool{},
				new(sync.Mutex),
			},
			args{
				"update player_cup t set t.cup=t.cup+1 where player_id=?",
				[]interface{}{1},
			},
			false,
		},
		{
			"function test",
			fields{
				CorrectMysqlTestUrl,
				10,
				map[*sql.DB]bool{},
				new(sync.Mutex),
			},
			args{
				"update player_cup t set t.cup=t.cup+1 where player_id=?",
				[]interface{}{1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				dsn:  tt.fields.dsn,
				size: tt.fields.size,
				dbs:  tt.fields.dbs,
				l:    tt.fields.l,
			}
			_, err := p.Query(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pool.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
