package gate

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// RegisterRoute : registered route
func (s *Service) RegisterRoute(router string, handler module.Handler) {
	if _, ok := s.routeMap.Load(router); ok {
		logger.Sugar.Fatal("repeated router : %s", router)
	}

	s.routeMap.Store(router, handler)
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
					args := make(map[string]interface{})
					// get `token`
					cookie, err := c.Request.Cookie("token")
					if err != nil {
						logger.Sugar.Infof("failed to get token : %v", err)
					}
					if cookie != nil {
						args[module.Token] = cookie.Value
					}
					// get request []bytes
					data, err := ioutil.ReadAll(c.Request.Body)
					if err != nil {
						logger.Sugar.Errorf("failed to get body : %v", err)
					}
					args[module.Request] = data

					resp := handler.Handler(args)
					logger.Sugar.Debugf("a post request for router [%s], response : %v", router, resp)
					// 目前使用protobuf作为通信协议
					// 由于gin框架支持protbuf，因此所有handler的resp都返回proto.Message,序列化由框架内部完成
					c.ProtoBuf(http.StatusOK, resp)
				})
			}
		case "GET":
			{
				s.gin.GET(router, func(c *gin.Context) {
					args := make(map[string]interface{})
					// get `token`
					cookie, err := c.Request.Cookie("token")
					if err != nil {
						logger.Sugar.Infof("failed to get token : %v", err)
					}
					if cookie != nil {
						args[module.Token] = cookie.Value
					}

					resp := handler.Handler(args)
					logger.Sugar.Debugf("a get request for router [%s], response : %v", router, resp)
					// 目前使用protobuf作为通信协议
					// 由于gin框架支持protbuf，因此所有handler的resp都返回proto.Message,序列化由框架内部完成
					c.ProtoBuf(http.StatusOK, resp)
				})
			}
		default:
			logger.Sugar.Fatalf("illegal http method [%s] for router [%s]", handler.Method, router)
		}
		return true
	})
}
