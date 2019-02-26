package config

import (
	"sync"

	"github.com/QHasaki/gp-server/module/v1"
	"github.com/QHasaki/gp-server/service/config/v1/source/github"
)

// Config describes Config model
type Config struct {
	sync.RWMutex
	configSource module.ConfigSource
	configMap    module.ConfigMap
}

// NewConfig is the constructor of config model
func NewConfig() *Config {
	c := &Config{
		configSource: source.NewGithub(),
	}

	c.InitConfig()

	return c
}
