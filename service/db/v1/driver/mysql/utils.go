package driver

import (
	"database/sql"

	"github.com/qinhan-shu/gp-server/logger"
)

// Resp describes the result of mysql query
type Resp struct {
	rows *sql.Rows
	done chan<- struct{}
}

// GetRows is to get query result
func (r *Resp) GetRows() *sql.Rows {
	return r.rows
}

// Close is to relesase resource
func (r *Resp) Close() {
	close(r.done)
}

// Query get rows from mysql db
func (p *MysqlDriver) Query(query string, args ...interface{}) (*Resp, error) {
	rows, err := p.conn.Query(query, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query rows: %v", err)
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		<-done
		rows.Close()
	}()

	return &Resp{
		rows: rows,
		done: done,
	}, nil
}

// QueryRow get at most one row from mysql db
func (p *MysqlDriver) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.conn.QueryRow(query, args...)
}

// Exec modify mysql db
func (p *MysqlDriver) Exec(sql string, args ...interface{}) error {
	res, err := p.conn.Exec(sql, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to update db: %v", err)
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		logger.Sugar.Errorf("failed to count affected rows: %v", err)
		return err
	}
	if count == 0 {
		return ErrNoRowsAffected
	}

	return nil
}
