package source_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/config/source/envs"
)

func TestGetConfig(t *testing.T) {
	gmSource, err := source.NewConfigSource()
	if err != nil {
		t.Error(err)
		return
	}

	config, err := gmSource.GetConfig()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", config)
}
