package gate

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// RegisterRoute : registered route
func (s *Service) RegisterRoute(router string, handler module.Router) {
	if _, ok := s.routeMap[router]; ok {
		logger.Sugar.Fatal("repeated router : %s", router)
	}

	s.routeMap[router] = handler
}

func (s *Service) registerToGate() {
	for router, handler := range s.routeMap {
		switch handler.Method {
		case "POST":
			{
				s.gin.POST(router, func(c *gin.Context) {
					logger.Sugar.Debugf("a post request for router [%s]", router)
					data, err := ioutil.ReadAll(c.Request.Body)
					if err != nil {
						logger.Sugar.Errorf("failed to get body : %v", err)
					}
					resp := handler.Handler(data)

					// 目前使用protobuf作为通信协议
					// 由于gin框架支持protbuf，因此所有handler的resp都返回proto.Message,序列化由框架内部完成
					c.ProtoBuf(http.StatusOK, resp)
				})
			}
		default:
			logger.Sugar.Fatalf("illegal http method [%s] for router [%s]", handler.Method, router)
		}
	}
}
