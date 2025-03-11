package redis

import (
	"time"
	"ztalk/internal/models"

	"github.com/go-redis/redis"
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
		Score:  float64(models.SecondsInOneDay),
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

// GetPostsScore
// 获取指定帖子的投票得分
// func GetPostsScore(ids []string) (data []float64, err error) {
// 	pipeline := rdb.Pipeline()
// 	cmds := make([]*redis.FloatCmd, len(ids))
// 	for i, id := range ids {
// 		key := KeyPostScoreZSet
// 		cmds[i] = pipeline.ZScore(key, id)
// 	}
// 	_, err = pipeline.Exec()
// 	if err != nil {
// 		return
// 	}
// 	for i, cmd := range cmds {
// 		if cmd.Err() != nil {
// 			err = fmt.Errorf("%s 分数信息错误：%v", ids[i], cmd.Err())
// 			return
// 		}
// 		data = append(data, cmd.Val())
// 	}
// 	return
// }

// GetSortedPostsInRange
// 根据参数获取帖子列表的ID
func GetSortedPostIDsInRange(pattern string, start, end int64) (res []string, err error) {
	var key string
	if pattern == models.OrderScore {
		key = KeyPostScoreZSet
	} else {
		key = KeyPostCreateTimeZSet
	}
	res, err = rdb.ZRevRange(key, start, end).Result()
	return
}
