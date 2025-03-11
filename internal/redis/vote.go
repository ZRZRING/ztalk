package redis

import (
	"ztalk/internal/models"
	"ztalk/pkg/message"
)

func GetVote(postID, userID string) (data float64, err error) {
	rdbKey := KeyVoteHash(postID)
	data, err = rdb.ZScore(rdbKey, userID).Result()
	return
}

func UpdateVote(postID, userID string, Value float64) (err error) {
	oldValue, _ := GetVote(postID, userID)
	delta := Value - oldValue
	if delta == 0 {
		return message.ErrVoteRepeated
	}
	pipeline := rdb.TxPipeline()
	delta *= models.ScorePerVote
	pipeline.ZIncrBy(KeyPostScoreZSet, delta, postID)
	if Value == 0 {
		pipeline.HDel(KeyVoteHash(postID), userID)
	} else {
		pipeline.HSet(KeyVoteHash(postID), userID, Value)
	}
	_, err = pipeline.Exec()
	return
}
