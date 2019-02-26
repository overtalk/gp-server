package config_test

import (
	"testing"

	"github.com/QHasaki/gp-server/service/config/v1"
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
