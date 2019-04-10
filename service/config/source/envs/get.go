package source

import (
	"fmt"
	"os"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// GetConfig return config
func (e *Env) GetConfig() (module.ConfigMap, error) {
	e.Lock()
	defer e.Unlock()

	conf := make(map[string]string)
	for _, key := range e.require {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			err := fmt.Errorf(`Config "%s" is absent`, key)
			logger.Sugar.Error(err)
			return nil, err
		}
		conf[key] = value
	}

	return conf, nil
}
