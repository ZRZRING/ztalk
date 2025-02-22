package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"ztalk/app/models"
	"ztalk/app/service"
	"ztalk/app/utils"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 判断 err 是不是 validator.ValidationErrors 类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)

		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "request param error",
			"error": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}
	fmt.Printf("注册用户: %v", p)
	if err := service.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}