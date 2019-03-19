package manage

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// authCheck : check weather the user is able to manager user
func authCheck(authCode int) bool {
	return authCode == int(protocol.Role_MANAGER)
}

// getReqAndToken : get token and protobuf data
func getReqAndToken(c *gin.Context) ([]byte, string, error) {
	// get data
	data, err := utils.GetRequestBody(c)
	if err != nil {
		return nil, "", err
	}
	// get token
	token, err := utils.GetToken(c)
	if err != nil {
		return nil, "", err
	}
	return data, token, nil
}

func getJudgeFileRelativePath(str string) string {
	node, _ := snowflake.NewNode(90)
	return "/" + node.Generate().String()
}
