package gate

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/test"
)

func addRoute(router *gin.Engine) {
	router.POST("/post", func(c *gin.Context) {
		result, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.Sugar.Errorf("failed to get body : %v", err)
		}

		req := &protocol.TestRequest{}
		if err := proto.Unmarshal(result, req); err != nil {
			logger.Sugar.Errorf("failed to unmarshal : %v", err)
			c.ProtoBuf(http.StatusOK, nil)
		} else {
			logger.Sugar.Infof("request from client : %v", req)
		}

		resp := test.Test(req)

		c.ProtoBuf(http.StatusOK, resp)
	})
}
