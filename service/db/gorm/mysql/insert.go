package db

// Insert : 插入数据，传入model即可（value）
// value 为struct的指针，插入成功之后，value中未赋值的字段将会被赋值
func (m *MysqlDriver) Insert(value interface{}) error {
	return m.conn.Create(value).Error
}
