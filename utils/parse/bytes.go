package parse

import (
	"github.com/qinhan-shu/gp-server/logger"
)

// Bytes turn ( interface{} ) to ( []byte )
func Bytes(in interface{}) []byte {
	var ret []byte

	switch in.(type) {
	case []byte:
		ret = in.([]byte)
	case string:
		ret = []byte(in.(string))
	case nil:
		return nil
	default:
		logger.Sugar.Error("parse to []byte error : unknown type ")
		return nil
	}

	return ret
}
