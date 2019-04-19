package judge

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/utils/judge"
)

func getJudgeConfig(language int64) *judge.LangConfig {
	var conf *judge.LangConfig
	switch language {
	case 1:
		{
			conf = judge.CLangConfig
		}
	case 2:
		{
			conf = judge.CPPLangConfig
		}
	case 3:
		{
			conf = judge.JavaLangConfig
		}
	case 4:
		{
			conf = judge.PY2LangConfig
		}
	default:
		// python3
		conf = judge.PY3LangConfig
	}
	return conf
}

func (j *Judge) getJudgeServer() (*judge.Client, error) {
	for _, conf := range j.judgeServer {
		client := judge.NewClient(judge.WithTimeout(0))
		client.SetOptions(judge.WithEndpointURL("http://"+conf.Addr), judge.WithToken(conf.Token))
		_, err := client.Ping()
		if err != nil {
			logger.Sugar.Errorf("ping response error : %v", err)
			continue
		}
		return client, nil
	}
	return nil, fmt.Errorf("no available judge server")
}
