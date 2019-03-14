package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

func checkDefaultValue(in interface{}) bool {
	switch in.(type) {
	case *model.User:
		return in.(*model.User).IsInited()
	default:
		return false
	}

}
