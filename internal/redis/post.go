package redis

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	oneDayInSeconds = 24 * 3600
)

// CreatePost
// 创建帖子
func CreatePost(id int64) (err error) {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(KeyPostCreateTimeZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	})
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  float64(oneDayInSeconds),
		Member: id,
	})
	_, err = pipeline.Exec()
	return
}

// GetPostCreateTime
// 获取指定帖子的创建时间
func GetPostCreateTime(id string) (data float64, err error) {
	data, err = rdb.ZScore(KeyPostCreateTimeZSet, id).Result()
	return
}

// GetPostScore
// 获取指定帖子的投票分数
func GetPostScore(id string) (data float64, err error) {
	data, err = rdb.ZScore(KeyPostScoreZSet, id).Result()
	return
}
