package mysql

import "ztalk/internal/models"

func InsertPost(p *models.Post) (err error) {
	sqlStr := `
insert into 
    post(post_id, title, content, author_id, community_id) 
values 
    (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(id int64) (data *models.Post, err error) {
	sqlStr := `
select 
    post_id, title, content, author_id, community_id 
from 
    post 
where 
    post_id = ?`
	data = new(models.Post)
	err = db.Get(data, sqlStr, id)
	return
}

func GetAllPosts(page, size int64) (data []*models.Post, err error) {
	sqlStr := `
select 
    post_id, title, content, author_id, community_id 
from 
    post 
ORDER BY 
    create_time DESC 
limit 
    ?, ?`
	data = make([]*models.Post, 0, size)
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}
