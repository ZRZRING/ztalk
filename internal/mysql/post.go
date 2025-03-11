package mysql

import (
	"strings"
	"ztalk/internal/models"

	"github.com/jmoiron/sqlx"
)

func InsertPost(p *models.Post) (err error) {
	sqlStr := `
		insert into post(post_id, title, content, author_id, community_id) 
		values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(id int64) (data *models.Post, err error) {
	sqlStr := `
		select post_id, title, content, author_id, community_id 
		from post 
		where post_id = ?`
	data = new(models.Post)
	err = db.Get(data, sqlStr, id)
	return
}

func GetAllPostsID(page, size int64) (data []int64, err error) {
	sqlStr := `
		select post_id 
		from post 
		ORDER BY create_time DESC 
		limit ?, ?`
	data = make([]int64, 0, size)
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}

func GetPostsByIDs(ids []string) (data []*models.Post, err error) {
	sqlStr := `
		select post_id, title, content, author_id, community_id, create_time
		from post
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	data = make([]*models.Post, 0, len(ids))
	err = db.Select(&data, query, args...)
	return
}
