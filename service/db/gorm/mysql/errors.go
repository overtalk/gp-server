package db

import (
	"errors"
)

var (
	// ErrNoMysqlConf : 缺少mysql配置
	ErrNoMysqlConf = errors.New("lack of mysql config")
)
