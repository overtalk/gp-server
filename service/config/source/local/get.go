package local

import (
	"encoding/json"
	"fmt"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// GetConfig return config
func (f *File) GetConfig() (module.ConfigMap, error) {
	const fileName = "server.json"

	data, err := f.fetch(fileName)
	if err != nil {
		logger.Sugar.Errorf("failed to get %s from config scorce (local file version) : %v", fileName, err)
		return nil, err
	}

	config := make(module.ConfigMap)

	if err := json.Unmarshal(data, &config); err != nil {
		logger.Sugar.Errorf("failed to decode %s: %v", fileName, err)
		return nil, err
	}

	if _, ok := config["ISSUCCEED"]; !ok {
		return nil, fmt.Errorf("failed to get config from config scorce (local file version), error message [ %v ]", config)
	}

	return config, nil
}
