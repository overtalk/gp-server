package parse

import (
	"errors"

	"github.com/QHasaki/Server/logger"
)

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
		logger.Sugar.Error(errors.New("parse to string error(unknown) : error type"))
		return nil
	}

	return ret
}
