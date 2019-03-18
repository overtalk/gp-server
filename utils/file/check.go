package file

import (
	"os"
)

// Exists : if the file/dir exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir : judge weather the path is a dir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile : judge weather the path is a file
func IsFile(path string) bool {
	return !IsDir(path)
}
