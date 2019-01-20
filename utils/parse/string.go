package parse

import (
	"strconv"

	"github.com/QHasaki/Server/logger"
)

func String(in interface{}) string {
	var ret string
	switch in.(type) {
	case string:
		ret = in.(string)
	case []uint8:
		ret = string(in.([]uint8))
	case int64:
		ret = strconv.FormatInt(in.(int64), 10)
	case int:
		ret = strconv.Itoa(in.(int))
	case nil:
		return ""
	default:
		logger.Sugar.Errorf("parse to string error(unknown) : %v", in)
		return ""
	}

	return ret
}
