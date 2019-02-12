package config

import (
	"fmt"

	"github.com/QHasaki/Server/logger"
)

// InitConfig is to get config
// Fatal if failed to get config
func (c *Config) InitConfig() {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		logger.Sugar.Fatalf("failed to init config : %v", err)
	}

	c.Lock()
	defer c.Unlock()

	c.configMap = configMap
}

// ReloadConfig is to reload config
// error if failed to get config
func (c *Config) ReloadConfig() error {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	c.configMap = configMap

	return nil
}

// GetConfigByConfigName is to get config value by config key
func (c *Config) GetConfigByConfigName(configName string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	configValue := c.configMap[configName]
	if configValue == "" {
		return "", fmt.Errorf("missing config : %s", configName)
	}

	return configValue, nil
}
