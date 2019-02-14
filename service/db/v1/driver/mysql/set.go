package driver

import (
	"strings"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Set is to modify db
// if where = nil, create a new record, or update the record
func (p *MysqlDriver) Set(document string, data model.Data, where model.Data) error {
	if where == nil {
		var (
			column []string
			needed []string
			values []interface{}
		)

		sql := "INSERT INTO `" + document + "` "

		for k, v := range data {
			column = append(column, k)
			needed = append(needed, "?")
			values = append(values, v)
		}

		resp, err := p.Query(sql+"("+strings.Join(column, ",")+") VALUES ("+strings.Join(needed, ", ")+")", values...)
		if err != nil {
			logger.Sugar.Errorf("failed to query row : %v", err)
			return err
		}

		defer resp.Close()
	} else {
		var (
			datas  []string
			wheres []string
			values []interface{}
		)

		sql := "UPDATE `" + document + "` SET "

		for k, v := range data {
			datas = append(datas, k+" = ?")
			values = append(values, v)
		}

		for k, v := range where {
			switch v.(type) {
			case model.Where:
				wheres = append(wheres, k+" "+v.(model.Where).Operator+" ?")
				values = append(values, v.(model.Where).Value)
			default:
				wheres = append(wheres, k+" = ?")
				values = append(values, v)
			}
		}

		resp, err := p.Query(sql+strings.Join(datas, ", ")+" WHERE "+strings.Join(wheres, " AND "), values...)
		if err != nil {
			logger.Sugar.Errorf("failed to query row : %v", err)
			return err
		}

		defer resp.Close()
	}
	return nil
}
