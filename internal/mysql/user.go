package mysql

import (
	"ztalk/internal/models"
)

func InsertUser(user *models.User) (err error) {
	sqlStr := `
insert into 
    user(user_id, username, password) 
values
    (?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func GetUserByUsername(username string) (user *models.User, err error) {
	sqlStr := `
select 
    user_id, username, password 
from 
    user 
where 
    username = ?`
	user = new(models.User)
	err = db.Get(user, sqlStr, username)
	return
}

func GetUserByID(id int64) (user *models.User, err error) {
	sqlStr := `
select 
    user_id, username, password 
from 
    user 
where 
    user_id = ?`
	user = new(models.User)
	err = db.Get(user, sqlStr, id)
	return
}
