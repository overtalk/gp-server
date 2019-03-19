package driver

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/qinhan-shu/gp-server/logger"
)

var (
	// ErrGetDBConn describes get db connection error
	ErrGetDBConn = errors.New("failed to get db connection")
	// ErrCloseMysqlPool describes close mysql pool error
	ErrCloseMysqlPool = errors.New("failed to close mysql pool")
	// ErrNoRowsAffected describes now rows affected error
	ErrNoRowsAffected = errors.New("no rows affected")
)

// MysqlPool describes a mysql connection pool
type Pool struct {
	dsn  string
	size int
	dbs  map[*sql.DB]bool
	l    *sync.Mutex
}

// NewMysqlPool creates a mysql connection pool
func NewMysqlPool(size int, addr, username, password, dbname string) (*Pool, error) {
	if size < 1 {
		logger.Sugar.Errorf("invalid mysql pool size: %d", size)
		return nil, errors.New("invalid mysql pool size")
	}
	pool := &Pool{
		dsn:  fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", username, password, addr, dbname),
		size: size,
		dbs:  make(map[*sql.DB]bool),
		l:    new(sync.Mutex),
	}

	return pool, nil
}

// CreateDBConn creates a new mysql db connection
func (p *Pool) CreateDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", p.dsn)
	if err != nil {
		logger.Sugar.Errorf("failed to open mysql: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logger.Sugar.Errorf("failed to ping mysql: %v", err)
		return nil, err
	}

	return db, nil
}

// Get a db connection from pool
func (p *Pool) Get() (*sql.DB, error) {
	p.l.Lock()
	defer p.l.Unlock()

	// check if any db connection is avaiable in the pool
	for db, ok := range p.dbs {
		if ok {
			// check if db connection is alive
			if err := db.Ping(); err != nil {
				// remove dead db connection
				delete(p.dbs, db)
			} else {
				p.dbs[db] = false
				return db, nil
			}
		}
	}

	// create new db connection if db pool is allowed to grow
	if len(p.dbs) < p.size {
		db, err := p.CreateDBConn()
		if err != nil {
			return nil, err
		}
		p.dbs[db] = false
		return db, nil
	}

	return nil, ErrGetDBConn
}

// Put a db connection into pool
func (p *Pool) Put(db *sql.DB) {
	p.l.Lock()
	defer p.l.Unlock()

	p.dbs[db] = true
}

// Close mysql connections
func (p *Pool) Close() error {
	p.l.Lock()
	defer p.l.Unlock()

	timeOut := time.After(10 * time.Second)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for len(p.dbs) > 0 {
		for db, ok := range p.dbs {
			if ok {
				delete(p.dbs, db)
			}
		}
		select {
		case <-timeOut:
			return ErrCloseMysqlPool
		case <-ticker.C:
		}
	}

	return nil
}
