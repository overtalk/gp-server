package local_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/config/source/local"
)

func TestGetConfig(t *testing.T) {
	gmSource := local.NewConfigSource()

	config, err := gmSource.GetConfig()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", config)
}
