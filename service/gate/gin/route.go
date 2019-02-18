package gate

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/protocol/v1"
)

func addRoute(router *gin.Engine) {

	router.POST("/post", func(c *gin.Context) {
		logger.Sugar.Info("POST")

		req := &protocol.TestRequest{}

		result, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.Sugar.Errorf("failed to get resp.body : %v", err)
		}

		if err := proto.Unmarshal(result, req); err != nil {
			logger.Sugar.Errorf("failed to unmarshal : %v", err)
		} else {
			logger.Sugar.Infof("request from client : %v", req)
		}

		// TODO: get resp, and marshal

		c.String(http.StatusOK, "Reload success")
	})
}
