package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"ztalk/app/models"
)

// CheckUserExists 检查用户是否存在
func CheckUserExists(username string) (err error) {
	sql := "select count(*) from users where username = ?"
	var count int
	if err = db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// encryptPassword 加密密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sql := "insert into users(user_id, username, password) values(?, ?, ?)"
	_, err = db.Exec(sql, user.UserID, user.Username, user.Password)
	return
}
