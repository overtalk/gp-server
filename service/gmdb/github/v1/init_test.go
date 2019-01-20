package gmdb_test

import (
	"fmt"
	"testing"

	"github.com/QHasaki/Server/service/gmdb/github/v1"
)

func TestFetch(t *testing.T) {
	gmSource := gmdb.NewGithub()
	serverByte, err := gmSource.Fetch("server.json")
	if err != nil {
		t.Errorf("failed to get server.json from gm scorce (github version) : %v", err)
		return
	}

	fmt.Println(string(serverByte))
}
