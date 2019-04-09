package local

import (
	"io/ioutil"
	"os"

	"github.com/qinhan-shu/gp-server/logger"
)

// File : local file config
type File struct {
	path string
}

// NewConfigSource : create config source
func NewConfigSource() *File {
	if os.Getenv("CONFIG_FILE_PATH") == "" {
		logger.Sugar.Fatal(`Environment "CONFIG_FILE_PATH" must be set`)
	}
	return &File{
		path: os.Getenv("CONFIG_FILE_PATH"),
	}
}

// Fetch is to get details from github
func (f *File) fetch(fileName string) ([]byte, error) {
	return ioutil.ReadFile(f.path + fileName)
}
