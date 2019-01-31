package configSource_test

import (
	"testing"

	"github.com/QHasaki/Server/service/configSource/github"
)

func TestFetch(t *testing.T) {
	gmSource := configSource.NewGithub()

	config, err := gmSource.GetConfig()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", config)
}
