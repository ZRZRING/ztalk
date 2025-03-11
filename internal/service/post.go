package service

import (
	"strconv"
	"ztalk/internal/models"
	"ztalk/internal/mysql"
	"ztalk/internal/redis"
	"ztalk/pkg/utils"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = utils.GenID()
	if err = mysql.InsertPost(p); err != nil {
		return
	}
	if err = redis.CreatePost(p.ID); err != nil {
		return
	}
	return
}

func GetPostDetail(id int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID() failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		return
	}
	data = &models.PostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return
}

func GetPosts(p *models.PostListParam) (data []*models.PostDetail, err error) {
	start, end := p.GetStartEnd()
	ids, err := redis.GetSortedPostIDsInRange(p.Order, start, end)
	if err != nil {
		zap.L().Error("redis.GetSortedPostsInRange() failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetSortedPostsInRange() return 0 data")
		return
	}
	data = make([]*models.PostDetail, 0, len(ids))
	for _, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			zap.L().Error("strconv.ParseInt() failed", zap.Error(err))
			continue
		}
		post, err := GetPostDetail(id)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err), zap.Any("postID", post))
			continue
		}
		data = append(data, post)
	}
	scores, err := redis.GetPostsScore(ids)
	if err != nil {
		zap.L().Error("redis.GetPostsVote() failed", zap.Error(err))
		return
	}
	for i, score := range scores {
		data[i].Score = int64(score)
	}
	return
}
