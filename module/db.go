package module

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// DB : 数据库模块接口
type DB interface {
	Test(id int) (*model.Test, error) // db模块用于测试的接口，后续进行删除
}
