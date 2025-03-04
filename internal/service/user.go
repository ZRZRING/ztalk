package service

import (
	"database/sql"
	"errors"
	"ztalk/internal/models"
	"ztalk/internal/mysql"
	"ztalk/pkg/jwt"
	"ztalk/pkg/message"
	"ztalk/pkg/utils"
)

// CheckUserExist 检查用户是否存在
func CheckUserExist(user *models.User) (err error) {
	userDb, err := mysql.GetUserByUsername(user.Username)
	if userDb != nil {
		return message.ErrUserExist
	}
	return
}

// CheckUserNotExist 检查用户是否存在
func CheckUserNotExist(user *models.User) (err error) {
	_, err = mysql.GetUserByUsername(user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return message.ErrUserNotExist
	}
	return
}

func SignUp(p *models.SignUpParam) (user *models.User, err error) {
	user = &models.User{
		UserID:   utils.GenID(),
		Username: p.Username,
		Password: utils.MD5(p.Password),
	}
	if err = CheckUserExist(user); err != nil {
		return
	}
	if err = mysql.InsertUser(user); err != nil {
		return
	}
	return
}

func Login(p *models.LoginParam) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: utils.MD5(p.Password),
	}
	if err = CheckUserNotExist(user); err != nil {
		return
	}
	userDb, err := mysql.GetUserByUsername(user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		err = message.ErrUserNotExist
		return
	}
	if userDb.Password != user.Password {
		err = message.ErrInvalidPassword
		return
	}
	user.UserID = userDb.UserID
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
