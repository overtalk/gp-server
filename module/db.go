package module

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// DB : database module
type DB interface {
	Test(id int) (*model.Test, error) // db模块用于测试的接口，后续进行删除
}
