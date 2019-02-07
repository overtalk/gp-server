package data

import (
	"github.com/QHasaki/Server/data/v1/driver/mysql"
)

// originDB defines db source
type originDB struct {
	conn *driver.Pool
}

// Set for originDB is to wirte db directly
func (p *originDB) Set(document string, data Data, where Data) error {
	if err := getCondition(document, data, where); err != nil {
		return err
	}
	return driver.Set(p.conn, document, data, where)
}

// Get for originDB is to read db directly
func (p *originDB) Get(document string, column []string, where Data) (Data, error) {
	if err := setCondition(document, column, where); err != nil {
		return nil, err
	}
	return driver.Get(p.conn, document, column, where)
}

func (p *originDB) Inc(document string, column []string, where Data) error {
	if err := setCondition(document, column, where); err != nil {
		return err
	}
	return driver.Inc(p.conn, document, column, where)
}
