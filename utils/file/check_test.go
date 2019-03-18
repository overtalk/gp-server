package file_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/utils/file"
)

func TestDir(t *testing.T) {
	path := "/Users/qinhan/go/src/github.com/qinhan-shu/mc/lru/list.go"
	if file.Exists(path) {
		if file.IsDir(path) {
			t.Logf("path[%s] is a dir", path)
			return
		}
		t.Logf("path[%s] is a file", path)
		return
	}
	t.Logf("path[%s] not exists", path)
}
