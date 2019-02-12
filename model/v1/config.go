package model

import (
	"errors"
)

var (
	// ErrInvalidConfigJSON defines invalid config json error
	ErrInvalidConfigJSON = errors.New("invalid config json")
)

// ConfigMap defines configurations for server
type ConfigMap map[string]string

// ConfigSource defines config source
type ConfigSource interface {
	GetConfig() (ConfigMap, error)
}

// Config defines config
type Config interface {
	InitConfig() error
	ReloadConfig() error
	GetConfigByConfigName(ConfigName string) (string, error)
}
