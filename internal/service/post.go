package service

import (
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

func GetPostList(offset, limit int64) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetAllPosts(offset, limit)
	if err != nil {
		return
	}
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err), zap.Any("postID", post))
			continue
		}
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err), zap.Any("postID", post))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName: user.Username,
			Post:       post,
			Community:  community,
		}
		data = append(data, postDetail)
	}
	return
}
