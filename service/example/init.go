package example

import (
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// Example : 业务逻辑模块的实现
type Example struct {
}

// NewExampleModule : Example模块的构造函数
func NewExampleModule() module.Example {
	return &Example{}
}

// Register : Example模块的注册函数，将改模块实现的功能注册到gate中
func Register(gate module.Gate) {
	exampleModule := NewExampleModule()
	gate.RegisterRoute("/example", module.Router{
		Method:  "POST",
		Handler: exampleModule.Example,
	})
}

// Example : Example模块的业务逻辑处理函数
func (t *Example) Example(data ...interface{}) interface{} {
	req := &protocol.TestRequest{}
	resp := &protocol.TestResponse{}
	if len(data) != 2 {
		resp.A = "no token"
		return resp
	}

	token := parse.String(data[0])
	logger.Sugar.Infof("get token : %s", token)

	if err := proto.Unmarshal(parse.Bytes(data[1]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.A = err.Error()
		return resp
	}

	logger.Sugar.Infof("request = %v", req)
	resp.A = "all ok"
	resp.B = "example"

	return resp
}
