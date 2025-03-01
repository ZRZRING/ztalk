package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ztalk/internal/models"
	"ztalk/internal/service"
	"ztalk/pkg/message"
	"ztalk/pkg/response"
	"ztalk/pkg/translate"
)

// checkValidator 判断 err 是不是 validator.ValidationErrors 类型，是则翻译成中文
func checkValidator(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		response.Error(c, response.CodeInvalidParam)
		return
	}
	msg := errs.Translate(translate.Trans)
	// 【强迫症】
	// msg = RemoveTopStruct(msg)
	response.ErrorWithMsg(c, response.CodeInvalidParam, msg)
	return
}

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		checkValidator(c, err)
		return
	}
	// zap.L().Info("SignUpParam", zap.Any("SignUpParam", p))
	user, err := service.SignUp(p)
	if err != nil {
		// zap.L().Error("service.SignUp() failed", zap.Error(err))
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
		// zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		checkValidator(c, err)
		return
	}
	// zap.L().Info("LoginParam", zap.Any("LoginParam", p))
	user, err := service.Login(p)
	if err != nil {
		// zap.L().Debug("username", zap.String("username", p.Username))
		// zap.L().Error("service.Login() failed", zap.Error(err))
		if errors.Is(err, message.ErrInvalidPassword) {
			response.Error(c, response.CodeInvalidPassword)
			return
		}
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, user)
}
