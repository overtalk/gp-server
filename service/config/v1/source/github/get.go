package source

import (
	"encoding/json"

	"github.com/QHasaki/gp-server/logger"
	"github.com/QHasaki/gp-server/module/v1"
)

// GetConfig return config
func (g *Github) GetConfig() (module.ConfigMap, error) {
	const fileName = "server.json"

	data, err := g.fetch("server.json")
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
		logger.Sugar.Error(ErrGetConfFail)
		return nil, ErrGetConfFail
	}

	return config, nil
}
