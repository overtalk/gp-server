package file

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/zip"
)

// Upload : upload file
func (f *File) Upload(r *http.Request) proto.Message {
	resp := &protocol.File{Status: &protocol.Status{}}

	// parse and validate file and post parameters
	targetFile, header, err := r.FormFile("uploadFile")
	if err != nil {
		logger.Sugar.Errorf("failed to get form file : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get form file"
		return resp
	}
	defer targetFile.Close()
	fileBytes, err := ioutil.ReadAll(targetFile)
	if err != nil {
		logger.Sugar.Errorf("failed to read file bytes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to read file bytes"
		return resp
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "application/zip":
		break
	default:
		logger.Sugar.Errorf("invalid file type : %s", filetype)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "invalid file type : " + filetype
		return resp
	}

	// create a dir to store zip
	results := strings.Split(header.Filename, ".")
	if len(results) != 2 || results[1] != "zip" {
		logger.Sugar.Errorf("invalid file type : %s", filetype)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "invalid file type : " + filetype
		return resp
	}

	dir := f.path + results[0]
	newPath := filepath.Join(dir, header.Filename)

	err = os.Mkdir(dir, os.ModePerm)
	if err != nil {
		logger.Sugar.Errorf("failed to make dir : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to make dir for judge file"
		return resp
	}

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		logger.Sugar.Errorf("failed to write zip file : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to write zip file"
		return resp
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		logger.Sugar.Errorf("failed to write zip file : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to write zip file"
		return resp
	}

	// unzip
	if err := zip.Unzip(dir+"/"+header.Filename, dir); err != nil {
		logger.Sugar.Errorf("failed to unzip : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to unzipe"
	}

	return resp
}
