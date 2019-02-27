package example

import (
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
)

// Example : 业务逻辑模块的实现
type Example struct {
	gate module.Gate
}

// NewExampleModule : Example模块的构造函数
func NewExampleModule(gate module.Gate) module.Example {
	return &Example{
		gate: gate,
	}
}

// Register : Example模块的注册函数，将改模块实现的功能注册到gate中
func (t *Example) Register() {
	t.gate.RegisterRoute("/example", module.Router{
		Method:  "POST",
		Handler: t.Example,
	})
}

// Example : Example模块的业务逻辑处理函数
func (t *Example) Example(data interface{}) interface{} {
	return &protocol.TestResponse{
		A: "111",
		B: "111",
	}
}
