package cache

import (
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

var (
	expired int64 = 30 * 60
	fileKey       = "FILE"
)

// SetFileItem : set the file item
func (r *RedisCache) SetFileItem(item *module.FileItem) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return err
	}

	if _, err := r.client.HSet(fileKey, item.ID, item.TS).Result(); err != nil {
		logger.Sugar.Errorf("failed to set file id : ", item.ID)
		return err
	}
	return nil
}

// DelFileItem : del file item
func (r *RedisCache) DelFileItem(id string) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return err
	}

	if _, err := r.client.HDel(fileKey, id).Result(); err != nil {
		logger.Sugar.Errorf("failed to del file id : ", id)
		return err
	}
	return nil
}

// GetExpiredFile : get experied file
func (r *RedisCache) GetExpiredFile() ([]string, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return nil, err
	}

	var files []string
	res := r.client.HGetAll(fileKey)
	if res.Err() != nil {
		logger.Sugar.Errorf("failed to get all keys in hash : %v", res.Err())
		return nil, res.Err()
	}
	result, err := res.Result()
	if err != nil {
		logger.Sugar.Errorf("failed to get all keys in hash : %v", err)
		return nil, err
	}

	now := time.Now().Unix()
	for key, value := range result {
		ts := parse.Int(value)
		if now-ts > expired {
			files = append(files, key)
		}
	}
	return files, nil
}
