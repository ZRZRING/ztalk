package service

import (
	"ztalk/internal/models"
	"ztalk/internal/mysql"
	"ztalk/pkg/utils"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = utils.GenID()
	if err = mysql.InsertPost(p); err != nil {
		return
	}
	return
}

func GetPostDetail(id int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return
	}
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
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
			continue
		}
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
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
