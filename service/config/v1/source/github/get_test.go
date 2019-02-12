package source_test

import (
	"testing"

	"github.com/QHasaki/Server/service/config/v1/source/github"
)

func TestFetch(t *testing.T) {
	gmSource := source.NewGithub()

	config, err := gmSource.GetConfig()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", config)
}
