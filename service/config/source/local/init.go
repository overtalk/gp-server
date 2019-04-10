package source

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qinhan-shu/gp-server/module"
)

// File : local file config
type File struct {
	path string
}

// NewConfigSource : create config source
func NewConfigSource() (module.ConfigSource, error) {
	if os.Getenv("CONFIG_FILE_PATH") == "" {
		return nil, fmt.Errorf(`Environment "CONFIG_FILE_PATH" must be set`)
	}
	return &File{
		path: os.Getenv("CONFIG_FILE_PATH"),
	}, nil
}

// Fetch is to get details from github
func (f *File) fetch(fileName string) ([]byte, error) {
	return ioutil.ReadFile(f.path + fileName)
}
