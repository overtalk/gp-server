package module

import (
	"errors"
)

var (
	// ErrInvalidConfigJSON : 配置文件不是JSON的错误
	ErrInvalidConfigJSON = errors.New("invalid config json")
)

// ConfigMap : 项目配置的map
type ConfigMap map[string]string

// ConfigSource : 配置文件数据源
type ConfigSource interface {
	GetConfig() (ConfigMap, error)
}

// Config : 配置模块接口
type Config interface {
	InitConfig() error
	ReloadConfig() error
	GetConfigByName(configName string) (string, error)
}
