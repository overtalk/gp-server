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
