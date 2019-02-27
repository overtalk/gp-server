package module

// Handler : 业务逻辑处理函数的格式
type Handler func(data interface{}) interface{}

// Router : http method & 业务逻辑函数
type Router struct {
	Method  string  // http方法
	Handler Handler // 处理函数（目前使用protobuf作为通信协议，返回的interface{}为proto.Message）
}
