package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils"
)

func TestRedisKey(t *testing.T) {
	key := utils.RedisKey("pvf", "result", "testID")
	if !assert.Equal(t, key, "pvf:result:testID") {
		t.Error("not expected")
		return
	}
}
