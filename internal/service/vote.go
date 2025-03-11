package service

import (
	"strconv"
	"time"
	"ztalk/internal/models"
	"ztalk/internal/redis"
	"ztalk/pkg/message"

	"go.uber.org/zap"
)

func Vote(id int64, p *models.VoteParam) (err error) {
	postID, userID := p.PostID, strconv.FormatInt(id, 10)

	// 判断投票时间限制
	postTime, err := redis.GetPostCreateTime(postID)
	if err != nil {
		zap.L().Error("redis.GetPostCreateTime() failed", zap.Error(err))
		return
	}
	period := float64(time.Now().Unix()) - postTime
	if period > models.SecondsInOneWeek {
		return message.ErrVoteTimeExpire
	}

	// 更新投票信息
	Value := float64(p.Direction)
	err = redis.UpdateVote(postID, userID, Value)
	if err != nil {
		zap.L().Error("redis.UpdateVote() failed", zap.Error(err))
		return
	}
	return
}
