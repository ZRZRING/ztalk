package mysql

import (
	"ztalk/internal/models"
)

func GetAllCommunities() (data []*models.Community, err error) {
	sqlStr := `
select 
    community_id, community_name, introduction, create_time 
from 
    community`
	err = db.Select(&data, sqlStr)
	return
}

func GetCommunityByID(id int64) (data *models.Community, err error) {
	sqlStr := `
select 
    community_id, community_name, introduction, create_time 
from 
    community 
where 
    community_id = ?`
	data = new(models.Community)
	err = db.Get(data, sqlStr, id)
	return
}
