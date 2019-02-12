package config

import (
	"sync"

	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/service/config/v1/source/github"
)

// Config describes Config model
type Config struct {
	sync.RWMutex
	configSource model.ConfigSource
	configMap    model.ConfigMap
}

// NewConfig is the constructor of config model
func NewConfig() *Config {
	c := &Config{
		configSource: source.NewGithub(),
	}

	c.InitConfig()

	return c
}
