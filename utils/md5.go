package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5 is to encrypt data using md5
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
