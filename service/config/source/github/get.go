package source

import (
	"encoding/json"
	"fmt"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// GetConfig return config
func (g *Github) GetConfig() (module.ConfigMap, error) {
	const fileName = "server.json"

	data, err := g.fetch(fileName)
	if err != nil {
		logger.Sugar.Errorf("failed to get %s from gm scorce (github version) : %v", fileName, err)
		return nil, err
	}

	config := make(module.ConfigMap)

	if err := json.Unmarshal(data, &config); err != nil {
		logger.Sugar.Errorf("failed to decode %s: %v", fileName, err)
		return nil, err
	}

	if _, ok := config["ISSUCCEED"]; !ok {
		return nil, ErrGetConfFail{
			info: fmt.Sprintf("%v", config),
		}
	}

	return config, nil
}
