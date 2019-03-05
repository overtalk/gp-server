package config_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/config"
)

func TestGetConfigByConfigName(t *testing.T) {
	c := config.NewConfig()

	configKey := "REDIS_ADDR"
	configValue, err := c.GetConfigByName(configKey)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s = %s", configKey, configValue)
}

func TestGetDataStorage(t *testing.T) {
	c := config.NewConfig()

	dataStorage, err := c.GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", dataStorage)

}
