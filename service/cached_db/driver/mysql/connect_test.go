package driver

import (
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		info DBInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"function test",
			args{
				DBInfo{
					Size:     10,
					Addr:     MysqlTestAddr,
					Username: "root",
					Password: "12341234",
					DBName:   "sausage_shooter",
				},
			},
			false,
		},
		{
			"function test2: size<1",
			args{
				DBInfo{
					Size:     0,
					Addr:     MysqlTestAddr,
					Username: "root",
					Password: "12341234",
					DBName:   "sausage_shooter",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
