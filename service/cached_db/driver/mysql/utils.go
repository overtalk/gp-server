package driver

import (
	"database/sql"

	"github.com/qinhan-shu/gp-server/logger"
)

type Resp struct {
	rows *sql.Rows
	done chan<- struct{}
}

func (r *Resp) GetRows() *sql.Rows {
	return r.rows
}

func (r *Resp) Close() {
	close(r.done)
}

func (p *Pool) Exec(sql string, args ...interface{}) error {
	db, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(db)

	res, err := db.Exec(sql, args...)
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

func (p *Pool) Query(query string, args ...interface{}) (*Resp, error) {
	db, err := p.Get()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		logger.Sugar.Errorf("failed to query rows: %v", err)
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		<-done
		rows.Close()
		p.Put(db)
	}()

	return &Resp{
		rows: rows,
		done: done,
	}, nil
}
