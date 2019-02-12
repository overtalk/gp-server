package config_test

import (
	"testing"

	"github.com/QHasaki/Server/service/config/v1"
)

func TestGetConfigByConfigName(t *testing.T) {
	c, err := config.NewConfig()
	if err != nil {
		t.Errorf("failed to new config : %v", err)
		return
	}
	configKey := "MYSQL_USERNAME"
	configValue, err := c.GetConfigByConfigName(configKey)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s = %s", configKey, configValue)
}
