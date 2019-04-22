package module

import (
	"errors"
)

var (
	// ErrInvalidConfigJSON : invalid json error
	ErrInvalidConfigJSON = errors.New("invalid config json")
)

// ConfigMap : configurations
type ConfigMap map[string]string

// ConfigSource : config data source
type ConfigSource interface {
	GetConfig() (ConfigMap, error)
}

// Config : configuration module
type Config interface {
	InitConfig()
	ReloadConfig() error
	GetConfigByName(configName string) (string, error)
	GetDataStorage() (*DataStorage, error)
	GetDataStorageConfigs() (*DataStorage, error)
	GetTestDB() (DB, error)
}
