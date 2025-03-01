package mysql

import (
	"ztalk/internal/models"
)

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user(user_id, username, password) values(?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (user *models.User, err error) {
	sqlStr := "select user_id, username, password from user where username = ?"
	user = new(models.User)
	err = db.Get(user, sqlStr, username)
	return
}

// GetUserById 根据用户ID获取用户
func GetUserById(userId int64) (user *models.User, err error) {
	sqlStr := "select user_id, username, password from user where user_id = ?"
	user = new(models.User)
	err = db.Get(user, sqlStr, userId)
	return
}
