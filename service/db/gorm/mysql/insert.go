package db

func (m *MysqlDriver) Insert(value interface{}) error {
	m.conn.Create(value)
	return nil
}
