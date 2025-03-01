package service

import (
	"ztalk/internal/models"
	"ztalk/internal/repository/mysql"
	"ztalk/pkg/jwt"
	utils2 "ztalk/pkg/utils"
)

func SignUp(p *models.SignUpParam) (user *models.User, err error) {
	user = &models.User{
		UserID:   utils2.GenID(),
		Username: p.Username,
		Password: utils2.MD5(p.Password),
	}
	err = mysql.Register(user)
	return
}

func Login(p *models.LoginParam) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: utils2.MD5(p.Password),
	}
	if err = mysql.Login(user); err != nil {
		return
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
