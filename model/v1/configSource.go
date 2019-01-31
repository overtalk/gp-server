package model

import (
	"errors"
)

var (
	// ErrInvalidConfigJSON defines invalid config json error
	ErrInvalidConfigJSON = errors.New("invalid config json")
)

// Config defines configurations
type ConfigMap map[string]string

// ConfigSource defines gm config source
type ConfigSource interface {
	GetConfig() (ConfigMap, error)
}
