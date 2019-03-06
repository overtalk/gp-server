package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/mode"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// GetToken is to get token key
func GetToken(playerID string) string {
	// in test mode, token = playerID
	if mode.GetMode() == mode.TestMode {
		return playerID
	}

	// return playerID + parse.String(time.Now().Unix()) + parse.String(rand.Int63()+rand.Int63())
	rand1, _ := utils.RandInt(0, 1000000)
	rand2, _ := utils.RandInt(0, 1000000)
	return playerID + "_" + parse.String(time.Now().Unix()) + "_" + parse.String(int64(rand1)+int64(rand2))
}
