package zip_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/utils/zip"
)

func TestUnzip1(t *testing.T) {
	var (
		path = "/Users/qinhan/go/src/github.com/qinhan-shu/gp-server/tmp/"
		file = "1.zip"
	)

	if err := zip.Unzip(path+file, path); err != nil {
		t.Error(err)
		return
	}
}
