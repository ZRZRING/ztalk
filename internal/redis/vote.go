package redis

import (
	"ztalk/pkg/message"
)

const (
	scorePerVote = oneDayInSeconds / 200
)

func GetVote(postID, userID string) (data float64, err error) {
	rdbKey := KeyVoteHash + postID
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
	pipeline.ZIncrBy(KeyPostScoreZSet, delta*scorePerVote, userID)
	KeyVoteHash := KeyVoteHash + postID
	if Value == 0 {
		pipeline.HDel(KeyVoteHash, userID)
	} else {
		pipeline.HSet(KeyVoteHash, userID, Value)
	}
	_, err = pipeline.Exec()
	return
}
