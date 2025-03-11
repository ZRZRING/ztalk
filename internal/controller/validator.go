package controller

import (
	"errors"
	"ztalk/pkg/response"
	"ztalk/pkg/translate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// checkValidator
// 判断 err 是不是 validator.ValidationErrors 类型，是则翻译成中文
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
