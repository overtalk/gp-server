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

// DataStorage defines data storage model
type DataStorage struct {
}

// ConfigSource defines gm config source
type ConfigSource interface {
	GetConfig() (ConfigMap, error)
}

// Config defines server config
type Config interface {
	GetDataStorage() (*DataStorage, error)
}
