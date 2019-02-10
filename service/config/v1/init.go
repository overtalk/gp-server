package config

import (
	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

// Config implemente Config interface in model
type Config struct {
	configSource model.ConfigSource
	configMap    model.ConfigMap
}

// NewConfig return Config
func NewConfig(configSource model.ConfigSource) *Config {
	c := &Config{
		configSource: configSource,
	}

	configMap, err := c.configSource.GetConfig()
	if err != nil {
		logger.Sugar.Fatalf("failed to get config : %v", err)
	}
	c.configMap = configMap

	return c
}

// Reload is to get config from config source
func (c *Config) Reload() {
	if c.configSource == nil {
		logger.Sugar.Errorf("[Reload Config Error] : nil configSource")
		return
	}
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		logger.Sugar.Fatalf("failed to get config : %v", err)
	}
	c.configMap = configMap
}
