package parse

import (
	"encoding/json"
	"strconv"

	"github.com/QHasaki/Server/logger"
)

// Int turn ( interface{} ) to ( int64 )
func Int(in interface{}) int64 {
	var ret int64
	jsonIn, ok := in.(json.Number)
	if ok {
		in = jsonIn.String()
	}
	switch in.(type) {
	case string:
		inp := in.(string)
		if inp == "" {
			return 0
		}
		left, err := strconv.ParseInt(inp, 10, 64)
		if err != nil {
			logger.Sugar.Errorf("parse to int error(string) : %v", err)
			return ret
		}
		ret = left
	case int:
		ret = int64(in.(int))
	case int32:
		ret = int64(in.(int32))
	case int64:
		ret = in.(int64)
	case uint:
		ret = int64(in.(uint))
	case uint16:
		ret = int64(in.(uint16))
	case uint32:
		ret = int64(in.(uint32))
	case uint64:
		ret = int64(in.(uint64))
	case float64:
		ret = int64(in.(float64))
	case nil:
		return 0
	default:
		logger.Sugar.Error("parse to int error : unknown type")
		return 0
	}

	return ret
}