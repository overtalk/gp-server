package judge

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/judge"
)

func getJudgeConfig(language protocol.Language) *judge.LangConfig {
	var conf *judge.LangConfig
	switch language {
	case protocol.Language_C:
		{
			conf = judge.CLangConfig
		}
	case protocol.Language_C_PLUS:
		{
			conf = judge.CPPLangConfig
		}
	case protocol.Language_JAVA:
		{
			conf = judge.JavaLangConfig
		}
	case protocol.Language_PYTHON3:
		{
			conf = judge.PY3LangConfig
		}
	default:
		// python2
		conf = judge.PY2LangConfig
	}
	return conf
}

func (j *Judge) getJudgeServer() (*judge.Client, error) {
	for _, conf := range j.judgeServer {
		client := judge.NewClient(judge.WithTimeout(0))
		client.SetOptions(judge.WithEndpointURL("http://"+conf.Addr), judge.WithToken(conf.Token))
		_, err := client.Ping()
		if err != nil {
			continue
		}
		return client, nil
	}
	return nil, fmt.Errorf("no available judge server")
}
