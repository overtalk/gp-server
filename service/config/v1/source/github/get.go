package source

import (
	"encoding/json"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// GetConfig return config
func (g *Github) GetConfig() (model.ConfigMap, error) {
	const fileName = "server.json"

	data, err := g.fetch("server.json")
	if err != nil {
		logger.Sugar.Errorf("failed to get %s from gm scorce (github version) : %v", fileName, err)
		return nil, err
	}

	config := make(model.ConfigMap)

	if err := json.Unmarshal(data, &config); err != nil {
		logger.Sugar.Errorf("failed to decode %s: %v", fileName, err)
		return nil, err
	}

	return config, nil
}
