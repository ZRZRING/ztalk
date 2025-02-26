package service

import (
	"ztalk/internal/models"
	"ztalk/internal/repository/mysql"
	"ztalk/internal/utils"
)

func SignUp(p *models.SignUpParam) (user *models.User, err error) {
	user = &models.User{
		UserID:   utils.GenID(),
		Username: p.Username,
		Password: utils.MD5(p.Password),
	}
	err = mysql.Register(user)
	return
}

func Login(p *models.LoginParam) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: utils.MD5(p.Password),
	}
	if err = mysql.Login(user); err != nil {
		return
	}
	token, err := utils.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
