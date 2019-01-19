package driver

import (
	"errors"
)

type DBInfo struct {
	Size 	 int
	Addr 	 string
	Username string
	Password string
	DBName 	 string
}

type Data = map[string]interface{}

type Where struct {
	Operator string
	Value 	 interface{}
}

func Connect(info DBInfo) (*Pool, error) {
	if info.Size < 1 {
		sugar.Errorf("invalid mysql pool size: %d", info.Size)
		return nil, errors.New("invalid mysql pool size")
	}
	pool, err := NewMysqlPool(info.Size, info.Addr, info.Username, info.Password, info.DBName)
	if err != nil {
		return nil, err
	}
	return pool, nil
}