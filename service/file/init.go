package file

import (
	"flag"

	"github.com/bwmarrin/snowflake"

	"github.com/qinhan-shu/gp-server/logger"
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
	node  *snowflake.Node
}

// NewFile : constructor for file
func NewFile(dataStorage *module.DataStorage) module.File {
	logger.Sugar.Infof("file path : %s", *uploadPath)
	node, err := snowflake.NewNode(10)
	if err != nil {
		logger.Sugar.Fatalf("failed to new node")
	}
	return &File{
		db:    dataStorage.DB,
		cache: dataStorage.Cache,
		path:  *uploadPath,
		node:  node,
	}
}

// Register : register module auth to gate
func Register(gate module.Gate, dataStorage *module.DataStorage) {
	module := NewFile(dataStorage)
	gate.RegisterRoute("/upload", "POST", module.Upload)
}
