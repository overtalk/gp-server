package manage

import (
	"github.com/bwmarrin/snowflake"

	"github.com/qinhan-shu/gp-server/protocol"
)

// authCheck : check weather the user is able to manager user
func authCheck(authCode int) bool {
	return authCode == int(protocol.Role_MANAGER)
}

func getJudgeFileRelativePath(str string) string {
	node, _ := snowflake.NewNode(90)
	return "/" + node.Generate().String()
}
