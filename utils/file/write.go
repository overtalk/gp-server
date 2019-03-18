package file

import (
	"fmt"
	"os"
)

// Write : write file
func Write(path string, writeBytes []byte) error {
	if Exists(path) {
		return fmt.Errorf("path[%s] is already exists", path)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write(writeBytes); err != nil {
		return err
	}
	return nil
}
