package file

import (
	"flag"
	"fmt"

	"github.com/qinhan-shu/gp-server/module"
)

var (
	uploadPath = flag.String("uploadPath", "./tmp", "upload path")
)

// File : implementation of judge module
type File struct {
	db    module.DB
	cache module.Cache
	path  string
}

// NewFile : constructor for file
func NewFile(dataStorage *module.DataStorage) module.File {
	fmt.Println("文件路径 ： ", *uploadPath)
	return &File{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
		path:  *uploadPath,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewFile(dataStorage)
	gate.RegisterRoute("/upload", "POST", module.Upload)
}
