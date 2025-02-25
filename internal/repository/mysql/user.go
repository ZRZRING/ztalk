package mysql

import (
	"database/sql"
	"errors"
	"ztalk/internal/models"
)

// CheckUserExist 检查用户是否存在
func CheckUserExist(user *models.User) (err error) {
	userDb, err := GetUserByUsername(user.Username)
	if userDb != nil {
		return ErrorUserExist
	}
	return
}

// CheckUserNotExist 检查用户是否存在
func CheckUserNotExist(user *models.User) (err error) {
	_, err = GetUserByUsername(user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	return
}

// Register 用户注册
func Register(user *models.User) (err error) {
	if err = CheckUserExist(user); err != nil {
		return
	}
	return InsertUser(user)
}

// Login 用户登录
func Login(user *models.User) (err error) {
	if err = CheckUserNotExist(user); err != nil {
		return
	}
	userDb, err := GetUserByUsername(user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	// zap.L().Info("提供密码", zap.String("password", userDb.Password))
	// zap.L().Info("正确密码", zap.String("password", user.Password))
	if userDb.Password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
