package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ztalk/internal/models"
	"ztalk/internal/repository/mysql"
	"ztalk/internal/service"
)

// checkValidator 判断 err 是不是 validator.ValidationErrors 类型，是则翻译成中文
func checkValidator(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		ResponseError(c, CodeInvalidParam)
		return
	}
	msg := errs.Translate(Trans)
	// 【强迫症】
	// msg = RemoveTopStruct(msg)
	ResponseErrorWithMsg(c, CodeInvalidParam, msg)
	return
}

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// zap.L().Error("注册时传入非法参数", zap.Error(err))
		checkValidator(c, err)
		return
	}
	// zap.L().Info("注册信息: ", zap.Any("SignUpParam", p))
	user, err := service.SignUp(p)
	if err != nil {
		// zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// zap.L().Error("登录时传入非法参数", zap.Error(err))
		checkValidator(c, err)
		return
	}
	// zap.L().Info("登录信息", zap.Any("LoginParam", p))
	user, err := service.Login(p)
	if err != nil {
		// zap.L().Error("登录失败", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
