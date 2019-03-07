package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils"
)

func TestRedisKey(t *testing.T) {
	key := utils.RedisKey("pvf", 1, "testID")
	expected := "pvf:1:testID"
	if !assert.Equal(t, key, expected) {
		t.Errorf("RedisKey[%s] not equal to expected[%s]", key, expected)
		return
	}
}
