package source

import (
	"sync"

	"github.com/qinhan-shu/gp-server/module"
)

// Env describes env config
type Env struct {
	sync.Mutex
	require []string
}

// NewConfigSource : create config source
func NewConfigSource() (module.ConfigSource, error) {
	r := []string{
		"REDIS_ADDR",
		"REDIS_PASS",
		"REDIS_POOLSIZE",
		"MYSQL_ADDR",
		"MYSQL_USER",
		"MYSQL_PASS",
		"MYSQL_DBNAME",
		"MYSQL_OPEN_CONNS_NUM",
		"MYSQL_IDLE_CONNS_NUM",
		"JUDGE_SERVER",
		// optional( debug | info | warn | error | dpanic | panic | fatal), default is info
		"LOG_LEVEL",
		"CERTFILE",
		"KEYFILE",
		"WEB_PORT",
		// for file server
		"FILE_PORT",
		"MAXUPLOADSIZE",
	}

	return &Env{
		require: r,
	}, nil
}
