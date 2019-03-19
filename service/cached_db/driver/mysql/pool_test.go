package driver

import (
	"database/sql"
	"reflect"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const CorrectMysqlTestUrl string = "root:12341234@tcp(localhost)/sausage_shooter?charset=utf8mb4&parseTime=true"
const InvalidMysqlTestUrl string = "root:1234123@tcp(localhost)/sausage_shooter?charset=utf8mb4&parseTime=true"

const MysqlTestAddr = "localhost"

func TestNewMysqlPool(t *testing.T) {
	type args struct {
		size     int
		addr     string
		username string
		password string
		dbname   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Pool
		wantErr bool
	}{
		{
			"function test:新建数据库连建池",
			args{
				size: 10,
				// addr:     "172.26.32.12",
				addr:     MysqlTestAddr,
				username: "root",
				password: "12341234",
				dbname:   "sausage_shooter",
			},
			&Pool{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			false,
		},
		{
			"function test:新建数据库连建池,大小小于1",
			args{
				size:     0,
				addr:     MysqlTestAddr,
				username: "sausage",
				password: "sausage_shooter",
				dbname:   "sausage_shooter",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMysqlPool(tt.args.size, tt.args.addr, tt.args.username, tt.args.password, tt.args.dbname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMysqlPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMysqlPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_CreateDBConn(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sql.DB
		wantErr bool
	}{
		{
			"function test：创建连接",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			nil,
			false,
		},
		{
			"function test：密码错误",
			fields{
				dsn:  InvalidMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			nil,
			true,
		},
		{
			"function test：dsn==nil",
			fields{
				dsn:  "",
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			nil,
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
			_, err := p.CreateDBConn()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pool.CreateDBConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPool_Get(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"function test：从连接池获取连接",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			false,
		},
		{
			"function test：获取的连接被关闭",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			false,
		},
		{
			"function test：pool.dsn错误,连接池无空闲连接",
			fields{
				dsn:  InvalidMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			true,
		},
		{
			"function test：连接池已满且无空闲连接",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			true,
		},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				dsn:  tt.fields.dsn,
				size: tt.fields.size,
				dbs:  tt.fields.dbs,
				l:    tt.fields.l,
			}
			switch index {
			case 0, 2:
				db, e := p.CreateDBConn()
				if e == nil {
					p.Put(db)
				}
			case 1:
				db, e := p.CreateDBConn()
				_ = db.Close()
				if e == nil {
					p.Put(db)
				}
			case 3:
				for i := 0; i < 10; i++ {
					db, e := p.CreateDBConn()
					if e == nil {
						p.dbs[db] = false
					}
				}
			}
			_, err := p.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pool.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPool_Put(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"function test：把连接放回连接池",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
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
			conn, _ := p.Get()
			p.Put(conn)
		})
	}
}

func TestPool_Close(t *testing.T) {
	type fields struct {
		dsn  string
		size int
		dbs  map[*sql.DB]bool
		l    *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"function test：关闭空连接池",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			false,
		},
		{
			"function test：关闭连接池",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			false,
		},
		{
			"function test：关闭无空闲连接的连接池",
			fields{
				dsn:  CorrectMysqlTestUrl,
				size: 10,
				dbs:  make(map[*sql.DB]bool),
				l:    new(sync.Mutex),
			},
			true,
		},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				dsn:  tt.fields.dsn,
				size: tt.fields.size,
				dbs:  tt.fields.dbs,
				l:    tt.fields.l,
			}
			switch index {
			case 1:
				for i := 0; i < 10; i++ {
					db, e := p.CreateDBConn()
					if e == nil {
						p.Put(db)
					}
				}
			case 2:
				for i := 0; i < 10; i++ {
					db, e := p.CreateDBConn()
					if e == nil {
						p.dbs[db] = false
					}
				}
			}
			if err := p.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Pool.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
