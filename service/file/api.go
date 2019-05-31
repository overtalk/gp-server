package file

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/zip"
)

// Upload : upload file
func (f *File) Upload(r *http.Request) proto.Message {
	resp := &protocol.File{Status: &protocol.Status{}}

	// parse and validate file and post parameters
	targetFile, _, err := r.FormFile("uploadFile")
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
	fileID := f.node.Generate().String()
	dir := f.path + fileID
	newFileName := fileID+".zip"
	newPath := filepath.Join(dir,newFileName)

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
	if err := zip.Unzip(dir+"/"+newFileName, dir); err != nil {
		logger.Sugar.Errorf("failed to unzip : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to unzip"
		return resp
	}

	if err := f.cache.SetFileItem(&module.FileItem{
		ID: fileID,
		TS: time.Now().Unix(),
	}); err != nil {
		logger.Sugar.Errorf("failed to set file info to redis : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to set file info to redis "
		return resp
	}

	resp.FileId = fileID

	logger.Sugar.Info("fileID = ", fileID)
	return resp
}
