package controller

import (
	"errors"
	"ztalk/internal/models"
	"ztalk/internal/service"
	"ztalk/pkg/message"
	"ztalk/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		checkValidator(c, err)
		return
	}
	user, err := service.SignUp(p)
	if err != nil {
		zap.L().Error("service.SignUp() failed", zap.Error(err))
		if errors.Is(err, message.ErrUserExist) {
			response.Error(c, response.CodeUserExist)
			return
		}
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, user)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		checkValidator(c, err)
		return
	}
	user, err := service.Login(p)
	if err != nil {
		zap.L().Error("service.Login() failed", zap.Error(err))
		if errors.Is(err, message.ErrInvalidPassword) {
			response.Error(c, response.CodeInvalidPassword)
			return
		}
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, user)
}
