package cache_test

import (
	"fmt"
	"testing"

	"github.com/qinhan-shu/gp-server/module"
)

func TestRedisCache_File(t *testing.T) {
	redisCache := getRedisCache(t)

	for i := 0; i < 100; i++ {
		if err := redisCache.SetFileItem(&module.FileItem{
			ID: fmt.Sprintf("key %d", i),
			TS: int64(i),
		}); err != nil {
			t.Error(err)
			return
		}
	}

	files, err := redisCache.GetExpiredFile()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(files)

	for _, file := range files {
		redisCache.DelFileItem(file)
	}

	t.Log("删除之后 : ")
	files, err = redisCache.GetExpiredFile()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(files)
}
