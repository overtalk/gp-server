package module

import (
	"github.com/golang/protobuf/proto"
)

// Handler : 业务逻辑处理函数的格式（protobuf版）
type Handler func(req proto.Message) proto.Message
