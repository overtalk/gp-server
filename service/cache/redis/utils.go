package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetToken is to get token key
func GetToken(userID string) string {
	// in test mode, token = userID
	if mode.GetMode() == mode.TestMode {
		return userID
	}

	// return userID + parse.String(time.Now().Unix()) + parse.String(rand.Int63()+rand.Int63())
	rand1, _ := utils.RandInt(0, 1000000)
	rand2, _ := utils.RandInt(0, 1000000)
	return userID + "_" + parse.String(time.Now().Unix()) + "_" + parse.String(int64(rand1)+int64(rand2))
}
