package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils"
)

func TestMD5(t *testing.T) {
	originStr := "qinhan"
	encryptedStr1 := utils.MD5(originStr)
	encryptedStr2 := utils.MD5(originStr)
	if !assert.Equal(t, encryptedStr1, encryptedStr2) {
		t.Errorf("the encryptedStr1[%s] is not equal to encryptedStr2[%s]", encryptedStr1, encryptedStr2)
		return
	}
}

func BenchmarkMD5(b *testing.B) {
	originStr := "xxxxx"
	for i := 0; i < b.N; i++ {
		utils.MD5(originStr)
	}
}
