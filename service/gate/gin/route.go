package gate

import (
	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// RegisterRoute : registered route
func (s *Service) RegisterRoute(router, method string, handler module.Handler) {
	if _, ok := s.routeMap.Load(router); ok {
		logger.Sugar.Fatal("repeated router : %s", router)
	}

	s.routeMap.Store(router, module.Router{
		Method:  method,
		Handler: handler,
	})
}

func (s *Service) registerToGate() {
	s.routeMap.Range(func(k, v interface{}) bool {
		router, err := parse.StringWithError(k)
		if err != nil {
			logger.Sugar.Fatalf("illegal http router[%v], not string, parse error [%v]", k, err)
		}
		handler := v.(module.Router)
		switch handler.Method {
		case "POST":
			{
				s.gin.POST(router, func(c *gin.Context) {
					code, resp := handler.Handler(c)
					logger.Sugar.Debugf("a post request for router [%s], response : %v", router, resp)
					// 目前使用protobuf作为通信协议
					// 由于gin框架支持protbuf，因此所有handler的resp都返回proto.Message,序列化由框架内部完成
					c.ProtoBuf(code, resp)
				})
			}
		case "GET":
			{
				s.gin.GET(router, func(c *gin.Context) {
					code, resp := handler.Handler(c)
					logger.Sugar.Debugf("a get request for router [%s], response : %v", router, resp)
					// 目前使用protobuf作为通信协议
					// 由于gin框架支持protbuf，因此所有handler的resp都返回proto.Message,序列化由框架内部完成
					c.ProtoBuf(code, resp)
				})
			}
		default:
			logger.Sugar.Fatalf("illegal http method [%s] for router [%s]", handler.Method, router)
		}
		return true
	})
}
