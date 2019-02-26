package source_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/service/config/v1/source/github"
)

func TestGetConfig(t *testing.T) {
	gmSource := source.NewGithub()

	config, err := gmSource.GetConfig()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", config)
}
