package module

// Handler : handler func format
type Handler func(data interface{}) interface{}

// Router : http method & handler func
type Router struct {
	Method  string  // http method (Get, Post ...)
	Handler Handler // handler func（use protobuf，return arg (interface{}) should be proto.Message）
}

// Gate : gateway module
type Gate interface {
	RegisterRoute(router string, handler Router)
}
