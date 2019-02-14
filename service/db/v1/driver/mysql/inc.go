package driver

import (
	"strings"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Inc is to add 1, used in some statistics
func (p *MysqlDriver) Inc(document string, columns []string, where model.Data) error {
	var (
		datas  []string
		column []string
		needed []string
		values []interface{}
	)

	sql := "INSERT INTO `" + document + "` "

	for k, v := range where {
		column = append(column, "`"+k+"`")
		needed = append(needed, "?")
		values = append(values, v)
	}

	for _, v := range columns {
		column = append(column, "`"+v+"`")
		needed = append(needed, "?")
		values = append(values, 1)
		datas = append(datas, "`"+v+"` = `"+v+"`+1")
	}

	resp, err := p.conn.Query(sql+"("+strings.Join(column, ",")+") VALUES ("+strings.Join(needed, ", ")+") ON DUPLICATE KEY UPDATE "+strings.Join(datas, ", "), values...)
	if err != nil {
		logger.Sugar.Errorf("failed to query row : %v", err)
		return err
	}

	defer resp.Close()
	return nil
}
