package cache

import (
	"github.com/go-redis/redis"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// SetRank set the rank
func (r *RedisCache) SetRank(item *module.RankItem) error {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return err
	}
	_, err := r.client.ZAdd(module.RankRedisKey,
		redis.Z{Score: float64(item.PassNum), Member: item.UserID}).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to SetRank of user %s: %v", item.UserID, err)
	}
	return err
}

// GetRank get the rank
func (r *RedisCache) GetRank() ([]*module.RankItem, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return nil, err
	}

	results, err := r.client.ZRevRangeWithScores(module.RankRedisKey, 0, 99).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to GetRank: %v", err)
		return nil, err
	}
	items := make([]*module.RankItem, 0, len(results))
	for _, v := range results {
		items = append(items, &module.RankItem{
			UserID:  parse.Int(v.Member),
			PassNum: int64(v.Score),
		})
	}
	return items, nil
}

// CleanRank clean some rank datas
func (r *RedisCache) CleanRank() (int64, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return 0, err
	}
	num, err := r.client.ZCard(module.RankRedisKey).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to ZCard redis: %v", err)
		return 0, err
	}
	availDeleteNum := num - module.MaxRanksNum
	if availDeleteNum <= 0 {
		return 0, nil
	}

	return r.client.ZRemRangeByRank(module.RankRedisKey, 0, availDeleteNum-1).Result()
}

// DelRank : del rank
func (r *RedisCache) DelRank() {
	r.client.Del(module.RankRedisKey)
}
