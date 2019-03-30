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
	_, err := r.client.ZAdd(rankRedisKey,
		redis.Z{Score: float64(item.PassNum), Member: item.UserID}).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to SetRank of user %s: %v", item.UserID, err)
	}
	return err
}

// GetRank get the rank
func (r *RedisCache) GetRank() ([]module.RankItem, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return nil, err
	}

	results, err := r.client.ZRevRangeWithScores(rankRedisKey, 0, 99).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to GetRank: %v", err)
		return nil, err
	}
	items := make([]module.RankItem, 0, len(results))
	for _, v := range results {
		items = append(items, module.RankItem{
			UserID:  parse.Int(v.Member),
			PassNum: int64(v.Score),
		})
	}
	return items, nil
}

// DelRank delete some rank datas
func (r *RedisCache) DelRank(deleteNum int64) (int64, error) {
	if _, err := r.client.Ping().Result(); err != nil {
		logger.Sugar.Errorf("failed to ping redis: %v", err)
		return 0, err
	}
	num, err := r.client.ZCard(rankRedisKey).Result()
	if err != nil {
		logger.Sugar.Errorf("failed to ZCard redis: %v", err)
		return 0, err
	}
	availDeleteNum := num - maxRanksNum
	if availDeleteNum <= 0 {
		return 0, nil
	}
	if deleteNum > availDeleteNum {
		deleteNum = availDeleteNum
	}
	return r.client.ZRemRangeByRank(rankRedisKey, 0, deleteNum-1).Result()
}
