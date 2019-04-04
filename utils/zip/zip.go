package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/qinhan-shu/gp-server/logger"
)

// UnZip : decode zip
func UnZip(fileName string, path string) error {
	r, err := zip.OpenReader(path + fileName)
	if err != nil {
		return err
	}

	for _, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			return fmt.Errorf("find dic in zip")
		}

		r, err := k.Open()
		if err != nil {
			logger.Sugar.Debug(err)
			continue
		}

		logger.Sugar.Debugf("unzip : %s", k.Name)
		defer r.Close()
		NewFile, err := os.Create(path + k.Name)
		if err != nil {
			logger.Sugar.Info(err)
			return err
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
	return nil
}
