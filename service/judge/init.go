package judge

import (
	"strings"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// Judge : implementation of judge module
type Judge struct {
	db          module.DB
	cache       module.Cache
	judgeServer []module.JudgeServerConf
}

// NewJudge : constructor for judge match
func NewJudge(dataStorage *module.DataStorage) module.Judge {
	value, isExist := dataStorage.Configs.Load("JUDGE_SERVER")
	if !isExist {
		logger.Sugar.Fatal(`config "JUDGE_SERVER" is absent`)
	}

	var judgeServer []module.JudgeServerConf
	addrs := strings.Split(value.(string), ",")
	for _, addr := range addrs {
		config := strings.Split(addr, "+")
		if len(config) != 2 {
			logger.Sugar.Fatal(`invalid format config "JUDGE_SERVER" : %s`, value.(string))
		}
		judgeServer = append(judgeServer, module.JudgeServerConf{
			Addr:  config[0],
			Token: config[1],
		})
	}

	logger.Sugar.Debugf("judge server : %v", judgeServer)

	return &Judge{
		db:          dataStorage.DB,
		cache:       dataStorage.Cache,
		judgeServer: judgeServer,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewJudge(dataStorage)

	gate.RegisterRoute("/judge", "POST", module.Judge)
}
