package service

import (
	"ztalk/app/models"
	"ztalk/app/repository/mysql"
	"ztalk/app/utils"
)

func SignUp(p *models.SignUpParam) (err error) {
	if err = mysql.CheckUserExists(p.Username); err != nil {
		return
	}
	userID := utils.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}
