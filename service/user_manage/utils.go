package user_manage

import (
	"github.com/qinhan-shu/gp-server/protocol"
)

// authCheck : check weather the user is able to manager user
func authCheck(authCode int) bool {
	return authCode == int(protocol.Role_MANAGER)
}
