package parse

import (
	"strconv"

	"github.com/qinhan-shu/gp-server/logger"
)

// String turn ( interface{} ) to ( string )
func String(in interface{}) string {
	out, _ := StringWithError(in)
	return out
}

// StringWithError turn ( interface{} ) to ( string ) with error
func StringWithError(in interface{}) (string, error) {
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
		return "", Err{method: "String", origin: in}
	default:
		logger.Sugar.Errorf("parse to string error(unknown) : %v", in)
		return "", Err{method: "String", origin: in}
	}

	return ret, nil
}
